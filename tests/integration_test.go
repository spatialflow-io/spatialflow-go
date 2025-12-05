//go:build integration

package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/spatialflow-io/spatialflow-go/spatialflow"
)

// Integration tests against local API
// Run with: go test -tags=integration ./tests/... -v

func getTestToken(t *testing.T) string {
	// Get a fresh token from the API
	baseURL := os.Getenv("SPATIALFLOW_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000"
	}

	loginPayload := []byte(`{"email":"test@spatialflow.io","password":"TestPassword123!"}`)
	resp, err := http.Post(baseURL+"/api/v1/auth/login", "application/json",
		bytes.NewReader(loginPayload))
	if err != nil {
		t.Fatalf("Failed to login: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("Login failed with status %d: %s", resp.StatusCode, string(body))
	}

	var loginResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("Failed to decode login response: %v", err)
	}

	return loginResp.AccessToken
}

func TestIntegration_HealthCheck(t *testing.T) {
	baseURL := os.Getenv("SPATIALFLOW_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000"
	}

	// Test without auth - health endpoint is public
	resp, err := http.Get(baseURL + "/api/v1/health")
	if err != nil {
		t.Fatalf("Health check failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Health check returned status %d, want 200", resp.StatusCode)
	}

	var health struct {
		Status  string `json:"status"`
		Service string `json:"service"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
		t.Fatalf("Failed to decode health response: %v", err)
	}

	if health.Status != "healthy" {
		t.Errorf("Health status = %q, want %q", health.Status, "healthy")
	}
}

func TestIntegration_ClientWithToken(t *testing.T) {
	baseURL := os.Getenv("SPATIALFLOW_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000"
	}

	token := getTestToken(t)

	client, err := spatialflow.NewClient(
		spatialflow.WithAccessToken(token),
		spatialflow.WithBaseURL(baseURL),
		spatialflow.WithTimeout(10*time.Second),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	ctx := context.Background()

	// Test listing geofences
	req, err := http.NewRequestWithContext(ctx, "GET", client.BaseURL()+"/api/v1/geofences", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	resp, err := client.HTTPClient().Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		t.Errorf("List geofences returned status %d: %s", resp.StatusCode, string(body))
	}

	// Verify response structure
	var listResp struct {
		Count   int           `json:"count"`
		Results []interface{} `json:"results"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&listResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	t.Logf("Found %d geofences", listResp.Count)
}

func TestIntegration_CheckResponse(t *testing.T) {
	baseURL := os.Getenv("SPATIALFLOW_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8000"
	}

	client, err := spatialflow.NewClient(
		spatialflow.WithAccessToken("invalid_token"),
		spatialflow.WithBaseURL(baseURL),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	ctx := context.Background()
	req, _ := http.NewRequestWithContext(ctx, "GET", client.BaseURL()+"/api/v1/geofences", nil)
	resp, err := client.HTTPClient().Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Should get 401 with invalid token
	if resp.StatusCode != 401 {
		t.Errorf("Expected 401, got %d", resp.StatusCode)
	}

	// CheckResponse should return an error
	if err := spatialflow.CheckResponse(resp); err == nil {
		t.Error("CheckResponse should return error for 401")
	}
}

func TestIntegration_Scheme_Host(t *testing.T) {
	client, err := spatialflow.NewClient(
		spatialflow.WithAccessToken("test"),
		spatialflow.WithBaseURL("http://localhost:8000"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client.Scheme() != "http" {
		t.Errorf("Scheme() = %q, want %q", client.Scheme(), "http")
	}

	if client.Host() != "localhost:8000" {
		t.Errorf("Host() = %q, want %q", client.Host(), "localhost:8000")
	}

	// Test HTTPS
	client2, _ := spatialflow.NewClient(
		spatialflow.WithAccessToken("test"),
		spatialflow.WithBaseURL("https://api.spatialflow.io"),
	)

	if client2.Scheme() != "https" {
		t.Errorf("Scheme() = %q, want %q", client2.Scheme(), "https")
	}

	if client2.Host() != "api.spatialflow.io" {
		t.Errorf("Host() = %q, want %q", client2.Host(), "api.spatialflow.io")
	}
}
