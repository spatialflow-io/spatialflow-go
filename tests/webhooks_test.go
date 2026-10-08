package tests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"github.com/spatialflow-io/spatialflow-go/v2/spatialflow"
)

// createSignature mirrors the backend: HMAC-SHA256 of the raw body, hex-encoded,
// prefixed with "sha256=".
func createSignature(payload, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return "sha256=" + hex.EncodeToString(mac.Sum(nil))
}

func TestVerifyWebhookSignature_Valid(t *testing.T) {
	secret := "whsec_test_secret"
	// Backend delivery shape: {"id", "event", "timestamp", "data"}.
	payload := `{"id":"del-1","event":"geofence.enter","timestamp":"2025-10-01T14:30:00Z","data":{"device_id":"dev123"}}`
	signature := createSignature(payload, secret)

	event, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0)
	if err != nil {
		t.Fatalf("verification failed: %v", err)
	}

	if event.Type != "geofence.enter" {
		t.Errorf("event.Type = %q, want %q (from backend \"event\" key)", event.Type, "geofence.enter")
	}
	if event.ID != "del-1" {
		t.Errorf("event.ID = %q, want %q", event.ID, "del-1")
	}
	if event.CreatedAt != "2025-10-01T14:30:00Z" {
		t.Errorf("event.CreatedAt = %q, want the backend timestamp", event.CreatedAt)
	}

	data, ok := event.Data["device_id"].(string)
	if !ok || data != "dev123" {
		t.Errorf("event.Data[device_id] = %v, want %q", event.Data["device_id"], "dev123")
	}
}

func TestVerifyWebhookSignature_LegacyTypePayload(t *testing.T) {
	secret := "whsec_test_secret"
	// Legacy/custom payload using "type"/"created_at" instead of "event"/"timestamp".
	payload := `{"type":"geofence.exit","created_at":"2025-10-01T15:00:00Z","data":{}}`
	signature := createSignature(payload, secret)

	event, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0)
	if err != nil {
		t.Fatalf("verification failed: %v", err)
	}
	if event.Type != "geofence.exit" {
		t.Errorf("event.Type = %q, want %q (legacy \"type\" key)", event.Type, "geofence.exit")
	}
	if event.CreatedAt != "2025-10-01T15:00:00Z" {
		t.Errorf("event.CreatedAt = %q, want the legacy created_at", event.CreatedAt)
	}
}

func TestVerifyWebhookSignature_BareHexNoPrefix(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	bareHex := hex.EncodeToString(mac.Sum(nil))

	if _, err := spatialflow.VerifyWebhookSignature([]byte(payload), bareHex, secret, 0); err != nil {
		t.Errorf("bare hex signature should verify: %v", err)
	}
}

func TestVerifyWebhookSignature_WrongSecret(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	signature := createSignature(payload, "wrong_secret")

	if _, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 0); err == nil {
		t.Error("expected verification to fail with wrong secret")
	}
}

func TestVerifyWebhookSignature_TamperedPayload(t *testing.T) {
	secret := "whsec_test_secret"
	signature := createSignature(`{"type":"original"}`, secret)

	if _, err := spatialflow.VerifyWebhookSignature([]byte(`{"type":"tampered"}`), signature, secret, 0); err == nil {
		t.Error("expected verification to fail with tampered payload")
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
		{"empty digest", "sha256="},
		{"invalid hex", "sha256=notvalidhex"},
		{"non-ascii", "sha256=éé"},
		{"old t=,v1= format", "t=1234567890,v1=abc123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := spatialflow.VerifyWebhookSignature([]byte(payload), tt.signature, secret, 0); err == nil {
				t.Errorf("expected verification to fail for %s", tt.name)
			}
		})
	}
}

func TestVerifyWebhookSignature_ToleranceIgnored(t *testing.T) {
	secret := "whsec_test_secret"
	payload := `{"type":"geofence.enter","data":{}}`
	signature := createSignature(payload, secret)

	// A non-zero tolerance must not change the outcome (deprecated/ignored).
	if _, err := spatialflow.VerifyWebhookSignature([]byte(payload), signature, secret, 1*time.Minute); err != nil {
		t.Errorf("tolerance should be ignored; verification should succeed: %v", err)
	}
}
