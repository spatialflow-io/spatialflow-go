package tests

import (
	"errors"
	"testing"
	"time"

	"github.com/spatialflow-io/spatialflow-go/spatialflow"
)

func TestNewClient_APIKey(t *testing.T) {
	client, err := spatialflow.NewClient(
		spatialflow.WithAPIKey("sf_test_key"),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}

	if client.BaseURL() != spatialflow.DefaultBaseURL {
		t.Errorf("BaseURL() = %q, want %q", client.BaseURL(), spatialflow.DefaultBaseURL)
	}
}

func TestNewClient_AccessToken(t *testing.T) {
	client, err := spatialflow.NewClient(
		spatialflow.WithAccessToken("eyJ..."),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client == nil {
		t.Fatal("NewClient() returned nil")
	}
}

func TestNewClient_CustomOptions(t *testing.T) {
	client, err := spatialflow.NewClient(
		spatialflow.WithAPIKey("sf_test"),
		spatialflow.WithBaseURL("http://localhost:8000"),
		spatialflow.WithTimeout(60*time.Second),
		spatialflow.WithMaxRetries(5),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	if client.BaseURL() != "http://localhost:8000" {
		t.Errorf("BaseURL() = %q, want %q", client.BaseURL(), "http://localhost:8000")
	}
}

func TestNewClient_NoAuth(t *testing.T) {
	_, err := spatialflow.NewClient()
	if err == nil {
		t.Error("NewClient() should fail without auth")
	}
}

func TestNewClient_BothAuthMethods(t *testing.T) {
	_, err := spatialflow.NewClient(
		spatialflow.WithAPIKey("sf_test"),
		spatialflow.WithAccessToken("eyJ..."),
	)
	if err == nil {
		t.Error("NewClient() should fail with both auth methods")
	}
}

func TestAPIError_Unwrap(t *testing.T) {
	apiErr := &spatialflow.APIError{
		StatusCode: 401,
		Message:    "Unauthorized",
		Err:        spatialflow.ErrAuthentication,
	}

	if !errors.Is(apiErr, spatialflow.ErrAuthentication) {
		t.Error("errors.Is(apiErr, ErrAuthentication) = false, want true")
	}
}

func TestAPIError_Error(t *testing.T) {
	tests := []struct {
		name    string
		err     *spatialflow.APIError
		want    string
	}{
		{
			name: "message only",
			err: &spatialflow.APIError{
				Message: "Not Found",
			},
			want: "Not Found",
		},
		{
			name: "message with detail",
			err: &spatialflow.APIError{
				Message: "Validation Error",
				Detail:  "name is required",
			},
			want: "Validation Error: name is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("Error() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSentinelErrors(t *testing.T) {
	// Verify all sentinel errors are distinct
	errs := []error{
		spatialflow.ErrAuthentication,
		spatialflow.ErrPermission,
		spatialflow.ErrNotFound,
		spatialflow.ErrValidation,
		spatialflow.ErrRateLimit,
		spatialflow.ErrServer,
		spatialflow.ErrTimeout,
		spatialflow.ErrConnection,
	}

	for i, err1 := range errs {
		for j, err2 := range errs {
			if i != j && errors.Is(err1, err2) {
				t.Errorf("sentinel errors %v and %v should not be equal", err1, err2)
			}
		}
	}
}
