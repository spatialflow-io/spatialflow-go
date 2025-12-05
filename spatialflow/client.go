package spatialflow

import (
	"net/http"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the default API base URL.
	DefaultBaseURL = "https://api.spatialflow.io"

	// DefaultTimeout is the default request timeout.
	DefaultTimeout = 30 * time.Second

	// DefaultMaxRetries is the default maximum retry attempts.
	DefaultMaxRetries = 3
)

// Client is the main SpatialFlow API client.
//
// After running code generation (./scripts/generate.sh go), the generated
// API client will be available. Use the raw HTTP client for direct API calls:
//
//	resp, err := client.HTTPClient().Get(client.BaseURL() + "/api/v1/geofences")
//
// Or use the helper methods for common operations like UploadGeofences.
type Client struct {
	config     *Config
	httpClient *http.Client

	// baseURLScheme and baseURLHost are parsed from BaseURL for generated client config
	baseURLScheme string
	baseURLHost   string
}

// NewClient creates a new SpatialFlow client with the given options.
//
// Example:
//
//	client, err := spatialflow.NewClient(
//	    spatialflow.WithAPIKey("sf_xxx"),
//	    spatialflow.WithTimeout(60 * time.Second),
//	)
func NewClient(opts ...Option) (*Client, error) {
	config := &Config{
		BaseURL:    DefaultBaseURL,
		Timeout:    DefaultTimeout,
		MaxRetries: DefaultMaxRetries,
	}

	for _, opt := range opts {
		opt(config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	// Create transport with auth and retry
	transport := &Transport{
		Base:       http.DefaultTransport,
		APIKey:     config.APIKey,
		Token:      config.AccessToken,
		MaxRetries: config.MaxRetries,
	}

	httpClient := config.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		}
	} else {
		// Wrap existing transport
		if httpClient.Transport != nil {
			transport.Base = httpClient.Transport
		}
		httpClient.Transport = transport
	}

	return &Client{
		config:        config,
		httpClient:    httpClient,
		baseURLScheme: extractScheme(config.BaseURL),
		baseURLHost:   stripScheme(config.BaseURL),
	}, nil
}

// Scheme returns the URL scheme (http or https) for the configured base URL.
// This is useful when configuring generated API clients.
func (c *Client) Scheme() string {
	return c.baseURLScheme
}

// Host returns the host (without scheme) for the configured base URL.
// This is useful when configuring generated API clients.
func (c *Client) Host() string {
	return c.baseURLHost
}

// HTTPClient returns the underlying HTTP client.
func (c *Client) HTTPClient() *http.Client {
	return c.httpClient
}

// BaseURL returns the configured base URL.
func (c *Client) BaseURL() string {
	return c.config.BaseURL
}

// stripScheme removes "https://" or "http://" from a URL, returns host:port.
func stripScheme(url string) string {
	if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}
	if strings.HasPrefix(url, "http://") {
		return strings.TrimPrefix(url, "http://")
	}
	return url
}

// extractScheme returns "https" or "http" from a URL, defaults to "https".
func extractScheme(url string) string {
	if strings.HasPrefix(url, "http://") {
		return "http"
	}
	return "https"
}
