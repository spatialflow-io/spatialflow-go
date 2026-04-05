// Package spatialflow provides the SpatialFlow Go SDK.
package spatialflow

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	// SDKVersion is the current version of the SDK.
	SDKVersion = "1.1.0"

	// UserAgent is the User-Agent header value.
	UserAgent = "spatialflow-go/" + SDKVersion
)

var (
	rngOnce sync.Once
	rng     *rand.Rand
)

func init() {
	// Seed PRNG once with current time for jitter
	rngOnce.Do(func() {
		rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
}

// retryableMethods are HTTP methods safe to retry (idempotent).
var retryableMethods = map[string]bool{
	http.MethodGet:     true,
	http.MethodHead:    true,
	http.MethodOptions: true,
	http.MethodPut:     true,
	http.MethodDelete:  true,
	http.MethodTrace:   true,
}

// Transport wraps http.RoundTripper to inject auth, User-Agent, and retry logic.
type Transport struct {
	// Base is the underlying transport. If nil, http.DefaultTransport is used.
	Base http.RoundTripper

	// APIKey is the API key for X-API-KEY header authentication.
	APIKey string

	// Token is the JWT token for Bearer token authentication.
	Token string

	// MaxRetries is the maximum number of retry attempts (default 3).
	MaxRetries int

	// BaseDelay is the initial retry delay (default 100ms).
	BaseDelay time.Duration

	// MaxDelay is the maximum retry delay (default 30s).
	MaxDelay time.Duration
}

// RoundTrip implements http.RoundTripper.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone request to avoid mutating original
	req = req.Clone(req.Context())

	// Inject User-Agent
	req.Header.Set("User-Agent", UserAgent)

	// Inject auth headers
	if t.APIKey != "" {
		req.Header.Set("X-API-KEY", t.APIKey)
	} else if t.Token != "" {
		req.Header.Set("Authorization", "Bearer "+t.Token)
	}

	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}

	// Determine if retryable
	maxRetries := t.MaxRetries
	if maxRetries == 0 {
		maxRetries = 3
	}

	// Don't retry non-idempotent methods
	if !retryableMethods[req.Method] {
		maxRetries = 0
	}

	// Don't retry if we have a body but can't recreate it
	// (GetBody is nil means we can't replay the request body)
	if req.Body != nil && req.GetBody == nil {
		maxRetries = 0
	}

	baseDelay := t.BaseDelay
	if baseDelay == 0 {
		baseDelay = 100 * time.Millisecond
	}
	maxDelay := t.MaxDelay
	if maxDelay == 0 {
		maxDelay = 30 * time.Second
	}

	var resp *http.Response
	var err error

	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Check context before each attempt
		if err := req.Context().Err(); err != nil {
			return nil, err
		}

		// Clone body for retry if needed
		if req.Body != nil && req.GetBody != nil && attempt > 0 {
			body, err := req.GetBody()
			if err != nil {
				return nil, err
			}
			req.Body = body
		}

		resp, err = base.RoundTrip(req)

		// Success or non-retryable error
		if err == nil && !shouldRetry(resp.StatusCode) {
			return resp, nil
		}

		// Don't retry if out of attempts
		if attempt >= maxRetries {
			break
		}

		// Calculate delay with exponential backoff + jitter
		delay := calculateBackoff(attempt, baseDelay, maxDelay)

		// Close response body before retry to avoid leaks
		if resp != nil {
			// Check Retry-After header for 429
			if resp.StatusCode == http.StatusTooManyRequests {
				if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
					if seconds, parseErr := strconv.Atoi(retryAfter); parseErr == nil {
						delay = time.Duration(seconds) * time.Second
					}
				}
			}
			resp.Body.Close()
		}

		// Wait with context cancellation support
		select {
		case <-time.After(delay):
			// Continue to next attempt
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}

	return resp, err
}

// shouldRetry returns true if the status code is retryable.
func shouldRetry(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests, // 429
		http.StatusInternalServerError,  // 500
		http.StatusBadGateway,           // 502
		http.StatusServiceUnavailable,   // 503
		http.StatusGatewayTimeout:       // 504
		return true
	default:
		return false
	}
}

// calculateBackoff calculates the backoff delay with exponential backoff and jitter.
func calculateBackoff(attempt int, baseDelay, maxDelay time.Duration) time.Duration {
	// Exponential backoff: baseDelay * 2^attempt
	delay := time.Duration(float64(baseDelay) * math.Pow(2, float64(attempt)))
	if delay > maxDelay {
		delay = maxDelay
	}
	// Add jitter (0-25%) using seeded PRNG
	jitter := time.Duration(rng.Float64() * 0.25 * float64(delay))
	return delay + jitter
}
