# SpatialFlow Go SDK

Official Go SDK for the [SpatialFlow](https://spatialflow.io) geospatial automation platform.

> **Alpha Release (v0.1.0)**: This SDK is in alpha. APIs may change before v1.0.

## Installation

```bash
go get github.com/spatialflow-io/spatialflow-go
```

**Requirements:** Go 1.21+

## Quick Start

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/spatialflow-io/spatialflow-go/spatialflow"
)

func main() {
    // Create client with API key
    client, err := spatialflow.NewClient(
        spatialflow.WithAPIKey("sf_your_api_key"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()

    // Make API requests using the HTTP client (includes auth, retry, User-Agent)
    req, _ := http.NewRequestWithContext(ctx, "GET", client.BaseURL()+"/api/v1/geofences", nil)
    resp, err := client.HTTPClient().Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    // Check for API errors
    if err := spatialflow.CheckResponse(resp); err != nil {
        log.Fatal(err)
    }

    log.Printf("Got %d response", resp.StatusCode)
}
```

## What's Included

This SDK provides:

| Feature | Status | Description |
|---------|--------|-------------|
| **HTTP Transport** | ✅ | Auth headers, User-Agent, retry with backoff |
| **Webhook Verification** | ✅ | HMAC-SHA256 signature validation |
| **File Uploads** | ✅ | Presigned URL workflow with job polling |
| **Pagination Helper** | ✅ | Generic `Pager[T]` for offset-based APIs |
| **Typed Errors** | ✅ | Sentinel errors with `errors.Is()` support |
| **Generated API Client** | 🔜 | Run `./scripts/generate.sh go` after setup |

The generated API client (for full CRUD on geofences, workflows, webhooks, devices) requires running code generation. See [Code Generation](#code-generation) below.

## Authentication

### API Key (Recommended)

```go
client, err := spatialflow.NewClient(
    spatialflow.WithAPIKey("sf_xxx"),
)
```

### JWT Access Token

```go
client, err := spatialflow.NewClient(
    spatialflow.WithAccessToken("eyJ..."),
)
```

## Configuration Options

```go
client, err := spatialflow.NewClient(
    spatialflow.WithAPIKey("sf_xxx"),
    spatialflow.WithBaseURL("https://api.spatialflow.io"),  // Default
    spatialflow.WithTimeout(60 * time.Second),               // Default: 30s
    spatialflow.WithMaxRetries(5),                           // Default: 3
)
```

## Making API Requests

The SDK provides an HTTP client with automatic auth, User-Agent, and retry logic:

```go
// The HTTPClient() includes auth headers and retry automatically
req, _ := http.NewRequestWithContext(ctx, "GET", client.BaseURL()+"/api/v1/geofences", nil)
resp, err := client.HTTPClient().Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()

// Check for API errors and get typed error
if err := spatialflow.CheckResponse(resp); err != nil {
    var apiErr *spatialflow.APIError
    if errors.As(err, &apiErr) {
        log.Printf("API error %d: %s", apiErr.StatusCode, apiErr.Detail)
    }
    return err
}

// Parse response...
```

## Error Handling

The SDK provides typed errors for common API error scenarios:

```go
import "errors"

resp, err := client.HTTPClient().Do(req)
if err != nil {
    return err
}

if err := spatialflow.CheckResponse(resp); err != nil {
    var apiErr *spatialflow.APIError
    if errors.As(err, &apiErr) {
        switch {
        case errors.Is(apiErr, spatialflow.ErrNotFound):
            log.Printf("Resource not found: %s", apiErr.Detail)
        case errors.Is(apiErr, spatialflow.ErrAuthentication):
            log.Printf("Auth failed: %s", apiErr.Detail)
        case errors.Is(apiErr, spatialflow.ErrRateLimit):
            log.Printf("Rate limited, retry after %v", apiErr.RetryAfter)
        case errors.Is(apiErr, spatialflow.ErrValidation):
            log.Printf("Validation error: %s", apiErr.Detail)
        default:
            log.Printf("API error: %s", apiErr.Error())
        }
    }
    return err
}
```

### Sentinel Errors

| Error | HTTP Status | Description |
|-------|-------------|-------------|
| `ErrAuthentication` | 401 | Invalid or missing credentials |
| `ErrPermission` | 403 | Insufficient permissions |
| `ErrNotFound` | 404 | Resource not found |
| `ErrValidation` | 400, 422 | Invalid request data |
| `ErrRateLimit` | 429 | Rate limit exceeded |
| `ErrServer` | 5xx | Server error |
| `ErrTimeout` | - | Request timeout |
| `ErrConnection` | - | Connection error |

## Webhook Verification

Verify incoming webhook signatures using HMAC-SHA256:

```go
import "github.com/spatialflow-io/spatialflow-go/spatialflow"

func handleWebhook(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    signature := r.Header.Get("X-SF-Signature")

    event, err := spatialflow.VerifyWebhookSignature(
        body,
        signature,
        webhookSecret,
        0, // Use default 5 minute tolerance
    )
    if err != nil {
        http.Error(w, "Invalid signature", http.StatusUnauthorized)
        return
    }

    switch event.Type {
    case "geofence.enter":
        handleGeofenceEnter(event.Data)
    case "geofence.exit":
        handleGeofenceExit(event.Data)
    }

    w.WriteHeader(http.StatusOK)
}
```

## File Uploads

Upload geofence files (GeoJSON, KML, GPX) with automatic job polling:

```go
result, err := client.UploadGeofences(ctx, spatialflow.UploadGeofencesOptions{
    FilePath:     "boundaries.geojson",
    GroupName:    "my-region",
    Timeout:      5 * time.Minute,
    PollInterval: 2 * time.Second,
    OnStatus: func(status string, resp *spatialflow.JobStatus) {
        log.Printf("Upload status: %s", status)
    },
})
if err != nil {
    return err
}

log.Printf("Created %d geofences (%d failed) in %v",
    result.CreatedCount, result.FailedCount, result.Duration)
```

## Pagination Helper

The SDK provides a generic `Pager` for paginated endpoints:

```go
// Create a pager with custom fetch function
pager := spatialflow.NewPager(100, func(ctx context.Context, offset, limit int) ([]Geofence, int, error) {
    // Fetch page from API using client.HTTPClient()
    // Return (items, totalCount, error)
    return fetchGeofences(ctx, offset, limit)
})

// Iterate page by page
for pager.HasMore() {
    items, err := pager.Next(ctx)
    if err != nil {
        return err
    }
    for _, item := range items {
        process(item)
    }
}

// Or collect all at once
all, err := pager.CollectAll(ctx)

// Or use channel-based iteration
for result := range pager.Iter(ctx) {
    if result.Err != nil {
        return result.Err
    }
    process(result.Item)
}
```

## Retry Logic

The SDK automatically retries failed requests with exponential backoff:

- **Retried status codes**: 429, 500, 502, 503, 504
- **Retried methods**: GET, HEAD, PUT, DELETE, OPTIONS (idempotent only)
- **Not retried**: POST, PATCH, or requests with non-replayable bodies
- **Backoff**: Exponential with 25% jitter
- **Rate limits**: Respects `Retry-After` header

Configure retry behavior:

```go
client, err := spatialflow.NewClient(
    spatialflow.WithAPIKey("sf_xxx"),
    spatialflow.WithMaxRetries(5),  // Default: 3
)
```

## Code Generation

To generate the full API client from the OpenAPI spec:

```bash
# From the monorepo root
cd sdks/scripts
./generate.sh go

# The generated client will be at:
# sdks/go/spatialflow/_generated/
```

After generation, you can configure the generated client with the SDK's HTTP client:

```go
import (
    "github.com/spatialflow-io/spatialflow-go/spatialflow"
    generated "github.com/spatialflow-io/spatialflow-go/spatialflow/_generated"
)

client, _ := spatialflow.NewClient(spatialflow.WithAPIKey("sf_xxx"))

// Configure generated client with SDK's HTTP client (includes auth/retry)
genConfig := generated.NewConfiguration()
genConfig.Scheme = client.Scheme()  // "https" or "http"
genConfig.Host = client.Host()      // "api.spatialflow.io"
genConfig.HTTPClient = client.HTTPClient()

genClient := generated.NewAPIClient(genConfig)

// Now use generated methods
geofences, resp, err := genClient.GeofencesApi.AppsGeofencesApiListGeofences(ctx).Execute()
```

## Local Development

```go
client, err := spatialflow.NewClient(
    spatialflow.WithAPIKey("sf_xxx"),
    spatialflow.WithBaseURL("http://localhost:8000"),
)
```

## Contributing

Contributions are welcome! Please see the [contributing guidelines](https://github.com/spatialflow-io/spatialflow-go/blob/main/CONTRIBUTING.md).

## License

MIT License - see [LICENSE](LICENSE) for details.
