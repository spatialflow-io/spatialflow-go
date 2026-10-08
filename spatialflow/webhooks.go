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

// DefaultTimestampTolerance is retained for backward compatibility with callers
// that passed it as the tolerance argument to VerifyWebhookSignature.
//
// Deprecated: tolerance is ignored because the SpatialFlow signature does not
// include a timestamp, so this constant has no effect.
const DefaultTimestampTolerance = 5 * time.Minute

// WebhookEvent represents a verified webhook payload.
//
// The backend delivers {"id", "event", "timestamp", "data"}: Type carries the
// backend "event" value and CreatedAt carries "timestamp". UnmarshalJSON also
// accepts the legacy/custom {"type", "created_at"} keys (e.g. from a custom
// payload template) so both shapes populate the struct.
type WebhookEvent struct {
	Type      string                 `json:"event"`
	Data      map[string]interface{} `json:"data"`
	ID        string                 `json:"id,omitempty"`
	CreatedAt string                 `json:"timestamp,omitempty"`
}

// UnmarshalJSON populates the event from the backend delivery shape
// ({"id","event","timestamp","data"}) while still accepting the legacy/custom
// {"type","created_at"} keys, keeping parity with the Node SDK so Go handlers
// switching on event.Type work for both payload shapes.
func (e *WebhookEvent) UnmarshalJSON(b []byte) error {
	var raw struct {
		ID        string                 `json:"id"`
		Event     string                 `json:"event"`
		Type      string                 `json:"type"`
		Timestamp string                 `json:"timestamp"`
		CreatedAt string                 `json:"created_at"`
		Data      map[string]interface{} `json:"data"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	e.ID = raw.ID
	e.Data = raw.Data
	e.Type = raw.Event
	if e.Type == "" {
		e.Type = raw.Type
	}
	e.CreatedAt = raw.Timestamp
	if e.CreatedAt == "" {
		e.CreatedAt = raw.CreatedAt
	}
	return nil
}

// VerifyWebhookSignature verifies the X-SF-Signature header.
//
// SpatialFlow signs each delivery with an HMAC-SHA256 of the raw request body
// and sends it in the X-SF-Signature header, hex-encoded and prefixed with
// "sha256=" (for example "sha256=5257a869..."). This function recomputes that
// HMAC and compares it in constant time.
//
// The tolerance parameter is deprecated and ignored: the SpatialFlow signature
// does not include a timestamp, so there is no time-based replay window. To
// guard against replays, deduplicate on the signed id in the payload,
// recording it in the same transaction as the work and acknowledging only
// committed work; the X-Idempotency-Key and X-SF-Event-ID headers are not
// signed. The default payload has an id; a custom payload template must
// include one.
//
// Example:
//
//	event, err := spatialflow.VerifyWebhookSignature(
//	    []byte(requestBody),
//	    request.Header.Get("X-SF-Signature"),
//	    webhookSecret,
//	    0, // tolerance is ignored
//	)
func VerifyWebhookSignature(payload []byte, signature, secret string, tolerance time.Duration) (*WebhookEvent, error) {
	_ = tolerance // deprecated and ignored; retained for API compatibility

	if signature == "" {
		return nil, errors.New("missing signature header")
	}

	// The header is "sha256=<hex>"; tolerate a bare hex digest as well.
	receivedSig := strings.TrimPrefix(strings.TrimSpace(signature), "sha256=")
	if receivedSig == "" {
		return nil, errors.New("invalid signature format: expected sha256=<hex digest>")
	}

	// Compute expected signature: HMAC-SHA256 of the raw request body.
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedBytes := mac.Sum(nil)

	// Decode received signature to bytes for constant-time comparison.
	receivedBytes, err := hex.DecodeString(receivedSig)
	if err != nil {
		return nil, errors.New("invalid signature: not valid hex")
	}

	// Constant-time comparison on raw bytes.
	if !hmac.Equal(receivedBytes, expectedBytes) {
		return nil, errors.New("signature verification failed")
	}

	// Parse and return event.
	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, fmt.Errorf("failed to parse webhook payload: %w", err)
	}

	return &event, nil
}

// DefaultWorkflowSignatureTolerance is the default maximum age, in either
// direction, of the X-SpatialFlow-Timestamp header.
const DefaultWorkflowSignatureTolerance = 5 * time.Minute

// VerifyWorkflowSignature verifies a workflow webhook action delivery and
// returns the parsed JSON body.
//
// A workflow Webhook action with a signing secret sends
// X-SpatialFlow-Timestamp (unix seconds) and X-SpatialFlow-Signature
// ("sha256=<hex>"), an HMAC-SHA256 of "<timestamp>.<raw body>". This is a
// different contract from the X-SF-Signature header handled by
// VerifyWebhookSignature.
//
// A tolerance of zero or less selects DefaultWorkflowSignatureTolerance. The
// workflow action sends whatever body the workflow configures, so a JSON body
// is decoded into an any (map[string]any for an object) without normalization,
// and any other body is returned as a string.
//
// Example:
//
//	body, err := spatialflow.VerifyWorkflowSignature(
//	    rawBody,
//	    r.Header.Get("X-SpatialFlow-Signature"),
//	    r.Header.Get("X-SpatialFlow-Timestamp"),
//	    webhookSecret,
//	    0, // default 5 minutes
//	)
func VerifyWorkflowSignature(payload []byte, signature, timestamp, secret string, tolerance time.Duration) (any, error) {
	if signature == "" {
		return nil, errors.New("missing signature header")
	}

	timestamp = strings.TrimSpace(timestamp)
	if timestamp == "" || strings.Trim(timestamp, "0123456789") != "" {
		return nil, errors.New("missing or invalid timestamp header")
	}
	sentAt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, errors.New("missing or invalid timestamp header")
	}

	if tolerance <= 0 {
		tolerance = DefaultWorkflowSignatureTolerance
	}
	age := time.Since(time.Unix(sentAt, 0))
	if age > tolerance || age < -tolerance {
		return nil, errors.New("timestamp outside the tolerance window")
	}

	receivedSig := strings.TrimPrefix(strings.TrimSpace(signature), "sha256=")
	if receivedSig == "" {
		return nil, errors.New("invalid signature format: expected sha256=<hex digest>")
	}
	receivedBytes, err := hex.DecodeString(receivedSig)
	if err != nil {
		return nil, errors.New("invalid signature: not valid hex")
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "."))
	mac.Write(payload)
	if !hmac.Equal(receivedBytes, mac.Sum(nil)) {
		return nil, errors.New("signature verification failed")
	}

	var body any
	if err := json.Unmarshal(payload, &body); err != nil {
		return string(payload), nil
	}
	return body, nil
}
