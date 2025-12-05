package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/spatialflow-io/spatialflow-go/spatialflow"
)

func TestTransport_AuthHeaders(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		token       string
		wantHeader  string
		wantValue   string
	}{
		{
			name:       "API key auth",
			apiKey:     "sf_test_key",
			wantHeader: "X-API-KEY",
			wantValue:  "sf_test_key",
		},
		{
			name:       "Bearer token auth",
			token:      "eyJ...",
			wantHeader: "Authorization",
			wantValue:  "Bearer eyJ...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var capturedReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				capturedReq = r
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			transport := &spatialflow.Transport{
				APIKey: tt.apiKey,
				Token:  tt.token,
			}
			client := &http.Client{Transport: transport}

			req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
			_, err := client.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}

			if got := capturedReq.Header.Get(tt.wantHeader); got != tt.wantValue {
				t.Errorf("header %s = %q, want %q", tt.wantHeader, got, tt.wantValue)
			}

			if got := capturedReq.Header.Get("User-Agent"); got != spatialflow.UserAgent {
				t.Errorf("User-Agent = %q, want %q", got, spatialflow.UserAgent)
			}
		})
	}
}

func TestTransport_Retry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &spatialflow.Transport{
		APIKey:     "sf_test",
		MaxRetries: 3,
		BaseDelay:  10 * time.Millisecond,
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	if attempts != 3 {
		t.Errorf("attempts = %d, want 3", attempts)
	}
}

func TestTransport_NoRetryOnPOST(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	transport := &spatialflow.Transport{
		APIKey:     "sf_test",
		MaxRetries: 3,
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodPost, server.URL, nil)
	resp, _ := client.Do(req)

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusInternalServerError)
	}

	if attempts != 1 {
		t.Errorf("attempts = %d, want 1 (no retry for POST)", attempts)
	}
}

func TestTransport_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &spatialflow.Transport{
		APIKey: "sf_test",
	}
	client := &http.Client{Transport: transport}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, server.URL, nil)
	_, err := client.Do(req)

	if err == nil {
		t.Error("expected context cancellation error")
	}
}

func TestTransport_RateLimitRetryAfter(t *testing.T) {
	attempts := 0
	start := time.Now()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	transport := &spatialflow.Transport{
		APIKey:     "sf_test",
		MaxRetries: 3,
		BaseDelay:  10 * time.Millisecond, // Would be much shorter without Retry-After
	}
	client := &http.Client{Transport: transport}

	req, _ := http.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := client.Do(req)
	elapsed := time.Since(start)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	// Should have waited at least 1 second due to Retry-After header
	if elapsed < 900*time.Millisecond {
		t.Errorf("elapsed = %v, expected at least ~1s due to Retry-After", elapsed)
	}
}
