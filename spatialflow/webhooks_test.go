package spatialflow

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"testing"
	"time"
)

const testWorkflowSecret = "workflow_secret" // pragma: allowlist secret

// signWorkflow mirrors the backend: HMAC-SHA256 over "<timestamp>.<body>".
func signWorkflow(body, timestamp, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + "." + body))
	return "sha256=" + hex.EncodeToString(mac.Sum(nil))
}

func unixAgo(d time.Duration) string {
	return strconv.FormatInt(time.Now().Add(-d).Unix(), 10)
}

func TestVerifyWorkflowSignatureGenuine(t *testing.T) {
	body := `{"alert":"geofence.enter","device":"d1"}`
	ts := unixAgo(0)

	got, err := VerifyWorkflowSignature([]byte(body), signWorkflow(body, ts, testWorkflowSecret), ts, testWorkflowSecret, 0)
	if err != nil {
		t.Fatalf("genuine delivery rejected: %v", err)
	}
	m, ok := got.(map[string]any)
	if !ok || m["alert"] != "geofence.enter" || m["device"] != "d1" {
		t.Fatalf("unexpected body: %#v", got)
	}
}

func TestVerifyWorkflowSignatureRejects(t *testing.T) {
	body := `{"a":1}`
	now := unixAgo(0)
	sub := hmac.New(sha256.New, []byte(testWorkflowSecret))
	sub.Write([]byte(body))
	subscriptionSig := "sha256=" + hex.EncodeToString(sub.Sum(nil))

	cases := []struct {
		name      string
		payload   string
		signature string
		timestamp string
		wantErr   string
	}{
		{"forged", body, signWorkflow(body, now, "wrong"), now, "verification failed"},
		{"tampered body", `{"a":2}`, signWorkflow(body, now, testWorkflowSecret), now, "verification failed"},
		{"subscription signature", body, subscriptionSig, now, "verification failed"},
		{"stale", body, signWorkflow(body, unixAgo(301*time.Second), testWorkflowSecret), unixAgo(301 * time.Second), "tolerance"},
		{"future", body, signWorkflow(body, unixAgo(-301*time.Second), testWorkflowSecret), unixAgo(-301 * time.Second), "tolerance"},
		{"far future", body, signWorkflow(body, "32503680000", testWorkflowSecret), "32503680000", "tolerance"},
		{"missing timestamp", body, signWorkflow(body, now, testWorkflowSecret), "", "timestamp"},
		{"garbage timestamp", body, signWorkflow(body, now, testWorkflowSecret), "abc", "timestamp"},
		{"decimal timestamp", body, signWorkflow(body, now, testWorkflowSecret), now + ".5", "timestamp"},
		{"negative timestamp", body, signWorkflow(body, now, testWorkflowSecret), "-5", "timestamp"},
		{"overflowing timestamp", body, signWorkflow(body, now, testWorkflowSecret), strings.Repeat("9", 30), "timestamp"},
		{"missing signature", body, "", now, "signature"},
		{"empty digest", body, "sha256=", now, "signature"},
		{"non hex digest", body, "sha256=zz", now, "signature"},
		{"changed timestamp", body, signWorkflow(body, now, testWorkflowSecret), unixAgo(-1 * time.Second), "verification failed"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := VerifyWorkflowSignature([]byte(tc.payload), tc.signature, tc.timestamp, testWorkflowSecret, 0)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("got err %v, want one containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestVerifyWorkflowSignatureCustomTolerance(t *testing.T) {
	body := `[1,2]`
	ts := unixAgo(10 * time.Minute)

	got, err := VerifyWorkflowSignature([]byte(body), signWorkflow(body, ts, testWorkflowSecret), ts, testWorkflowSecret, 15*time.Minute)
	if err != nil {
		t.Fatalf("rejected within custom tolerance: %v", err)
	}
	if arr, ok := got.([]any); !ok || len(arr) != 2 {
		t.Fatalf("unexpected body: %#v", got)
	}
}

func TestVerifyWebhookSignatureUnchanged(t *testing.T) {
	body := `{"id":"d1","event":"webhook.test","timestamp":"2025-10-01T14:30:00Z","data":{}}`
	mac := hmac.New(sha256.New, []byte("test_secret_key"))
	mac.Write([]byte(body))
	sig := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	event, err := VerifyWebhookSignature([]byte(body), sig, "test_secret_key", 0)
	if err != nil {
		t.Fatalf("subscription delivery rejected: %v", err)
	}
	if event.Type != "webhook.test" || event.ID != "d1" {
		t.Fatalf("unexpected event: %#v", event)
	}
}

func TestVerifyWorkflowSignatureReturnsTextBody(t *testing.T) {
	now := strconv.FormatInt(time.Now().Unix(), 10)
	body := "alert=entered&zone=dock"
	got, err := VerifyWorkflowSignature([]byte(body), signWorkflow(body, now, testWorkflowSecret), now, testWorkflowSecret, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != body {
		t.Fatalf("got %#v, want the text body", got)
	}
}
