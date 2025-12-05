package tests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/spatialflow-io/spatialflow-go/spatialflow"
)

func createSignature(payload, secret string, timestamp int64) string {
	signedPayload := fmt.Sprintf("%d.%s", timestamp, payload)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(signedPayload))
	sig := hex.EncodeToString(mac.Sum(nil))
	return fmt.Sprintf("t=%d,v1=%s", timestamp, sig)
}

func TestVerifyWebhookSignature_Valid(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{"device_id":"dev123"}}`
	timestamp := time.Now().Unix()
	signature := createSignature(payload, secret, timestamp)

	event, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0)
	if err != nil {
		t.Fatalf("verification failed: %v", err)
	}

	if event.Type != "geofence.enter" {
		t.Errorf("event.Type = %q, want %q", event.Type, "geofence.enter")
	}

	data, ok := event.Data["device_id"].(string)
	if !ok || data != "dev123" {
		t.Errorf("event.Data[device_id] = %v, want %q", event.Data["device_id"], "dev123")
	}
}

func TestVerifyWebhookSignature_InvalidSignature(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	timestamp := time.Now().Unix()

	// Create signature with wrong secret
	signature := createSignature(payload, "wrong_secret", timestamp)

	_, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0)
	if err == nil {
		t.Error("expected verification to fail with wrong secret")
	}
}

func TestVerifyWebhookSignature_ExpiredTimestamp(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	// 10 minutes ago (past default 5 minute tolerance)
	timestamp := time.Now().Add(-10 * time.Minute).Unix()
	signature := createSignature(payload, secret, timestamp)

	_, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0)
	if err == nil {
		t.Error("expected verification to fail with expired timestamp")
	}
}

func TestVerifyWebhookSignature_FutureTimestamp(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	// 10 minutes in the future (past default 5 minute tolerance)
	timestamp := time.Now().Add(10 * time.Minute).Unix()
	signature := createSignature(payload, secret, timestamp)

	_, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0)
	if err == nil {
		t.Error("expected verification to fail with future timestamp")
	}
}

func TestVerifyWebhookSignature_InvalidFormat(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`

	tests := []struct {
		name      string
		signature string
	}{
		{"empty signature", ""},
		{"missing timestamp", "v1=abc123"},
		{"missing signature", "t=1234567890"},
		{"invalid timestamp", "t=notanumber,v1=abc123"},
		{"invalid hex", "t=1234567890,v1=notvalidhex"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := spatialflow.VerifyWebhookSignature([]byte(payload), tt.signature, secret, 0)
			if err == nil {
				t.Errorf("expected verification to fail for %s", tt.name)
			}
		})
	}
}

func TestVerifyWebhookSignature_CustomTolerance(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	// 3 minutes ago
	timestamp := time.Now().Add(-3 * time.Minute).Unix()
	signature := createSignature(payload, secret, timestamp)

	// Should fail with 1 minute tolerance
	_, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 1*time.Minute)
	if err == nil {
		t.Error("expected verification to fail with 1 minute tolerance")
	}

	// Should succeed with 5 minute tolerance
	_, err = spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 5*time.Minute)
	if err != nil {
		t.Errorf("verification should succeed with 5 minute tolerance: %v", err)
	}
}
