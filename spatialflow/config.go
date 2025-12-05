package spatialflow

import (
	"errors"
	"net/http"
	"time"
)

// Config holds client configuration.
type Config struct {
	// APIKey is the API key for authentication (X-API-KEY header).
	APIKey string

	// AccessToken is the JWT token for authentication (Bearer header).
	AccessToken string

	// BaseURL is the API base URL.
	BaseURL string

	// Timeout is the request timeout.
	Timeout time.Duration

	// MaxRetries is the maximum number of retry attempts.
	MaxRetries int

	// HTTPClient is a custom HTTP client. If set, Transport will wrap its transport.
	HTTPClient *http.Client
}

// validate checks that the config is valid.
func (c *Config) validate() error {
	if c.APIKey != "" && c.AccessToken != "" {
		return errors.New("provide either APIKey or AccessToken, not both")
	}
	if c.APIKey == "" && c.AccessToken == "" {
		return errors.New("APIKey or AccessToken is required")
	}
	return nil
}

// Option configures the client.
type Option func(*Config)

// WithAPIKey sets the API key for authentication.
func WithAPIKey(key string) Option {
	return func(c *Config) { c.APIKey = key }
}

// WithAccessToken sets the JWT access token.
func WithAccessToken(token string) Option {
	return func(c *Config) { c.AccessToken = token }
}

// WithBaseURL sets a custom API base URL.
// Example: "https://api.spatialflow.io" or "http://localhost:8000"
func WithBaseURL(url string) Option {
	return func(c *Config) { c.BaseURL = url }
}

// WithHTTPClient sets a custom HTTP client.
// Note: The client's transport will be wrapped with auth and retry logic.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Config) { c.HTTPClient = client }
}

// WithTimeout sets the request timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Config) { c.Timeout = d }
}

// WithMaxRetries sets the maximum number of retry attempts.
func WithMaxRetries(n int) Option {
	return func(c *Config) { c.MaxRetries = n }
}
