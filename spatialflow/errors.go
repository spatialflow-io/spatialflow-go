package spatialflow

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// Sentinel errors for error checking with errors.Is().
var (
	ErrAuthentication = errors.New("authentication failed")
	ErrPermission     = errors.New("permission denied")
	ErrNotFound       = errors.New("resource not found")
	ErrValidation     = errors.New("validation failed")
	ErrConflict       = errors.New("resource conflict")
	ErrRateLimit      = errors.New("rate limit exceeded")
	ErrServer         = errors.New("server error")
	ErrTimeout        = errors.New("request timeout")
	ErrConnection     = errors.New("connection error")
)

// APIError represents an API error response.
type APIError struct {
	// StatusCode is the HTTP status code.
	StatusCode int

	// Message is the high-level error message.
	Message string

	// Detail is the detailed error message from the API.
	Detail string

	// ErrorCode is the machine-readable error code from the API.
	ErrorCode string

	// RetryAfter is the duration to wait before retrying (for rate limit errors).
	RetryAfter time.Duration

	// Err is the underlying sentinel error for errors.Is() support.
	Err error
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Detail)
	}
	return e.Message
}

// Unwrap returns the underlying sentinel error for errors.Is() support.
func (e *APIError) Unwrap() error {
	return e.Err
}

// errorResponse matches the API error format.
type errorResponse struct {
	Detail    string                 `json:"detail"`
	ErrorCode string                 `json:"error_code,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// NewAPIError creates an APIError from an HTTP response.
func NewAPIError(resp *http.Response) *APIError {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
		Message:    http.StatusText(resp.StatusCode),
	}

	// Parse response body for detail
	if resp.Body != nil {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		if len(body) > 0 {
			var errResp errorResponse
			if json.Unmarshal(body, &errResp) == nil {
				apiErr.Detail = errResp.Detail
				apiErr.ErrorCode = errResp.ErrorCode
			}
		}
	}

	// Parse Retry-After for 429
	if resp.StatusCode == http.StatusTooManyRequests {
		if retryAfter := resp.Header.Get("Retry-After"); retryAfter != "" {
			if seconds, err := strconv.Atoi(retryAfter); err == nil {
				apiErr.RetryAfter = time.Duration(seconds) * time.Second
			}
		}
	}

	// Map status code to sentinel error
	switch resp.StatusCode {
	case http.StatusUnauthorized:
		apiErr.Err = ErrAuthentication
	case http.StatusForbidden:
		apiErr.Err = ErrPermission
	case http.StatusNotFound:
		apiErr.Err = ErrNotFound
	case http.StatusBadRequest, http.StatusUnprocessableEntity:
		apiErr.Err = ErrValidation
	case http.StatusConflict:
		apiErr.Err = ErrConflict
	case http.StatusTooManyRequests:
		apiErr.Err = ErrRateLimit
	default:
		if resp.StatusCode >= 500 {
			apiErr.Err = ErrServer
		}
	}

	return apiErr
}

// CheckResponse returns an error if the response indicates failure.
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return NewAPIError(resp)
}

// WrapGenericError converts generic openapi-generator errors to typed APIError.
// Use this when the generated client returns an error.
func WrapGenericError(err error, resp *http.Response) error {
	if err == nil {
		return nil
	}
	if resp != nil && resp.StatusCode >= 400 {
		return NewAPIError(resp)
	}
	// Connection/timeout errors
	return &APIError{
		Message: err.Error(),
		Err:     ErrConnection,
	}
}
