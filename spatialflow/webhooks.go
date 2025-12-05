package spatialflow

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DefaultTimestampTolerance is 5 minutes (past and future).
const DefaultTimestampTolerance = 5 * time.Minute

// WebhookEvent represents a verified webhook payload.
type WebhookEvent struct {
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
	ID        string                 `json:"id,omitempty"`
	CreatedAt string                 `json:"created_at,omitempty"`
}

// VerifyWebhookSignature verifies the X-SF-Signature header.
// Format: t=<unix_timestamp>,v1=<hmac_sha256_hex>
//
// The tolerance parameter applies to both past and future timestamps
// to handle clock skew. Use 0 for DefaultTimestampTolerance (5 min).
//
// Example:
//
//	event, err := spatialflow.VerifyWebhookSignature(
//	    []byte(requestBody),
//	    request.Header.Get("X-SF-Signature"),
//	    webhookSecret,
//	    0, // use default tolerance
//	)
func VerifyWebhookSignature(payload []byte, signature, secret string, tolerance time.Duration) (*WebhookEvent, error) {
	if tolerance == 0 {
		tolerance = DefaultTimestampTolerance
	}

	// Parse signature header
	parts := strings.Split(signature, ",")
	var timestamp int64
	var receivedSig string
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "t=") {
			ts, err := strconv.ParseInt(strings.TrimPrefix(part, "t="), 10, 64)
			if err != nil {
				return nil, errors.New("invalid timestamp in signature")
			}
			timestamp = ts
		} else if strings.HasPrefix(part, "v1=") {
			receivedSig = strings.TrimPrefix(part, "v1=")
		}
	}

	if timestamp == 0 || receivedSig == "" {
		return nil, errors.New("invalid signature format: missing t= or v1=")
	}

	// Check timestamp freshness (both past and future)
	signedAt := time.Unix(timestamp, 0)
	age := time.Since(signedAt)
	if age > tolerance {
		return nil, fmt.Errorf("signature expired: signed %v ago (tolerance: %v)", age, tolerance)
	}
	if age < -tolerance {
		return nil, fmt.Errorf("signature from future: %v ahead (tolerance: %v)", -age, tolerance)
	}

	// Compute expected signature: HMAC-SHA256(secret, "timestamp.payload")
	signedPayload := fmt.Sprintf("%d.%s", timestamp, string(payload))
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signedPayload))
	expectedBytes := mac.Sum(nil)

	// Decode received signature to bytes for constant-time comparison
	receivedBytes, err := hex.DecodeString(receivedSig)
	if err != nil {
		return nil, errors.New("invalid signature: not valid hex")
	}

	// Constant-time comparison on raw bytes
	if !hmac.Equal(receivedBytes, expectedBytes) {
		return nil, errors.New("signature verification failed")
	}

	// Parse and return event
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, fmt.Errorf("failed to parse webhook payload: %w", err)
	}

	return &event, nil
}
