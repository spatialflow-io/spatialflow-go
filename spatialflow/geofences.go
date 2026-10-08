package spatialflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// BulkItemInput is a single address item in a bulk-create request.
type BulkItemInput struct {
	// Address is the street address to geocode and fence (required, max 500 chars).
	Address string `json:"address"`

	// Name is an optional display name (max 255 chars). If omitted, derived from address.
	Name *string `json:"name,omitempty"`

	// BufferMeters is the fence radius in meters (1-10000, default 100).
	BufferMeters *int `json:"buffer_meters,omitempty"`

	// Tags is an optional list of tag strings (max 20 per item).
	Tags []string `json:"tags,omitempty"`
}

// BulkCreateRequest is the request body for POST /api/v1/geofences/bulk.
type BulkCreateRequest struct {
	// Items is the list of address items to geocode and create as geofences.
	Items []BulkItemInput `json:"items"`

	// DedupStrategy controls duplicate handling:
	//   "skip"     (default) — duplicates surface as status "duplicate" with dedup_match info.
	//   "override" — duplicates are created anyway; each result carries dedup_match for awareness.
	//   "fail"     — first duplicate aborts the batch; HTTP 409 returned.
	DedupStrategy string `json:"dedup_strategy,omitempty"`
}

// BulkResultError is the per-item error when status is "error".
type BulkResultError struct {
	// Message is a human-readable description of the failure.
	Message string `json:"message"`

	// Code is the machine-readable error code:
	//   GEOCODE_FAILED, GEOCODER_RATE_LIMIT, INVALID_TAGS, CREATE_FAILED.
	Code string `json:"code"`
}

// BulkDedupMatch describes the existing geofence matched during dedup.
type BulkDedupMatch struct {
	// GeofenceID is the ID of the existing geofence.
	GeofenceID string `json:"geofence_id"`

	// Name is the name of the existing geofence.
	Name string `json:"name"`

	// DistanceMeters is the distance from the geocoded point to the existing fence centroid.
	// Pointer type because the backend returns null for address-path matches (no spatial distance computed).
	DistanceMeters *float64 `json:"distance_meters,omitempty"`
}

// BulkResult is the per-item result in a BulkCreateResponse.
type BulkResult struct {
	// Index is the 0-based position of this item in the request Items slice.
	Index int `json:"index"`

	// Status is one of "created", "duplicate", or "error".
	Status string `json:"status"`

	// GeofenceID is the UUID of the newly created geofence (only when status is "created").
	GeofenceID *string `json:"geofence_id,omitempty"`

	// Error contains error details (only when status is "error").
	Error *BulkResultError `json:"error,omitempty"`

	// ErrorCode mirrors Error.Code as the stable top-level error code.
	ErrorCode *string `json:"error_code,omitempty"`

	// DedupMatch describes the existing geofence that triggered dedup detection
	// (present when status is "duplicate", or on "created" with dedup_strategy="override").
	DedupMatch *BulkDedupMatch `json:"dedup_match,omitempty"`
}

// BulkCreateResponse is the response from POST /api/v1/geofences/bulk.
type BulkCreateResponse struct {
	// Results is the per-item result list in the same order as the request Items slice.
	Results []BulkResult `json:"results"`
}

// BulkCreateGeofences creates geofences from a list of address items.
//
// This is the integration entry point for syncing address-based geofences
// from CRM exports, TMS stop lists, or CSV imports. Per-item atomicity:
// a single item failure does not roll back previous items.
//
// Example (dedup-override pattern — detect then force):
//
//	resp, err := client.BulkCreateGeofences(ctx, spatialflow.BulkCreateRequest{
//	    Items: []spatialflow.BulkItemInput{
//	        {Address: "123 Main St, SF, CA", Tags: []string{"customer:acme"}},
//	        {Address: "456 Market St, SF, CA"},
//	    },
//	    DedupStrategy: "skip",
//	})
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, r := range resp.Results {
//	    switch r.Status {
//	    case "created":
//	        fmt.Printf("Created: %s\n", *r.GeofenceID)
//	    case "duplicate":
//	        fmt.Printf("Duplicate at index %d (existing: %s)\n", r.Index, r.DedupMatch.GeofenceID)
//	    case "error":
//	        fmt.Printf("Error at index %d: %s (%s)\n", r.Index, r.Error.Message, r.Error.Code)
//	    }
//	}
func (c *Client) BulkCreateGeofences(ctx context.Context, req BulkCreateRequest) (*BulkCreateResponse, error) {
	if req.DedupStrategy == "" {
		req.DedupStrategy = "skip"
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.config.BaseURL+"/api/v1/geofences/bulk",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	var result BulkCreateResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}
