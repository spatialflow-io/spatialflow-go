package spatialflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// UploadGeofencesOptions configures a geofence file upload.
type UploadGeofencesOptions struct {
	// FilePath is the path to the file to upload.
	FilePath string

	// GroupName is an optional group name for the imported geofences.
	GroupName string

	// Timeout is the total timeout for upload + polling (default 5m).
	Timeout time.Duration

	// PollInterval is the interval between status checks (default 2s).
	PollInterval time.Duration

	// OnStatus is called after each status check during polling.
	OnStatus func(status string, response *JobStatus)
}

// PresignedURLRequest is the request body for getting a presigned URL.
type PresignedURLRequest struct {
	FileType string `json:"file_type"`
	Filename string `json:"filename"`
	FileSize int64  `json:"file_size"`
}

// PresignedURLResponse is the response from the presigned URL endpoint.
type PresignedURLResponse struct {
	UploadURL   string `json:"upload_url"`
	Key         string `json:"key"`
	FileID      string `json:"file_id"`
	ExpiresIn   int    `json:"expires_in"`
	ContentType string `json:"content_type"`
}

// UploadGeofencesRequest is the request body for starting a geofence upload job.
type UploadGeofencesRequest struct {
	FileID    string  `json:"file_id"`
	GroupName *string `json:"group_name,omitempty"`
}

// UploadGeofencesResponse is the response from starting a geofence upload job.
type UploadGeofencesResponse struct {
	JobID     string `json:"job_id"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	StatusURL string `json:"status_url"`
}

// UploadGeofences uploads a geofence file and polls for completion.
//
// Workflow:
// 1. POST /api/v1/storage/presigned-url → get upload_url + file_id
// 2. PUT file to presigned URL (S3)
// 3. POST /api/v1/geofences/upload with file_id → get job_id
// 4. Poll GET /api/v1/geofences/upload/{job_id}/status until complete
//
// Example:
//
//	result, err := client.UploadGeofences(ctx, spatialflow.UploadGeofencesOptions{
//	    FilePath:  "boundaries.geojson",
//	    GroupName: "my-region",
//	})
func (c *Client) UploadGeofences(ctx context.Context, opts UploadGeofencesOptions) (*JobResult, error) {
	// Defaults
	if opts.Timeout == 0 {
		opts.Timeout = 5 * time.Minute
	}
	if opts.PollInterval == 0 {
		opts.PollInterval = 2 * time.Second
	}

	// Create timeout context
	ctx, cancel := context.WithTimeout(ctx, opts.Timeout)
	defer cancel()

	// Validate and open file
	file, err := os.Open(opts.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat file: %w", err)
	}

	// Determine content type from extension (for S3 upload)
	ext := filepath.Ext(opts.FilePath)
	contentType := contentTypeFromExt(ext)

	if contentType == "" {
		return nil, fmt.Errorf("unsupported file type: %s (supported: .geojson, .json, .kml, .gpx)", ext)
	}

	// Step 1: Get presigned upload URL
	// Note: file_type is always "geofences" (not per-format like "geojson")
	// This matches Python/Node SDKs and the actual API contract
	presignReq := PresignedURLRequest{
		FileType: "geofences",
		Filename: stat.Name(),
		FileSize: stat.Size(),
	}

	presignResp, err := c.getPresignedURL(ctx, presignReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get presigned URL: %w", err)
	}

	// Step 2: Upload file to S3 (streaming)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, presignResp.UploadURL, file)
	if err != nil {
		return nil, fmt.Errorf("failed to create upload request: %w", err)
	}
	req.Header.Set("Content-Type", contentType)
	req.ContentLength = stat.Size()

	// Use a plain HTTP client (no auth needed for S3 presigned URLs)
	s3Client := &http.Client{Timeout: opts.Timeout}
	resp, err := s3Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}
	defer resp.Body.Close()

	// S3 returns 200 or 204 on success
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("upload failed with status %d", resp.StatusCode)
	}

	// Step 3: Start import job
	var groupName *string
	if opts.GroupName != "" {
		groupName = &opts.GroupName
	}

	uploadReq := UploadGeofencesRequest{
		FileID:    presignResp.FileID,
		GroupName: groupName,
	}

	uploadResp, err := c.startGeofenceUpload(ctx, uploadReq)
	if err != nil {
		return nil, fmt.Errorf("failed to start import: %w", err)
	}

	// Step 4: Poll for completion
	return c.PollJob(ctx, PollJobOptions{
		JobID: uploadResp.JobID,
		FetchStatus: func(ctx context.Context) (*JobStatus, error) {
			return c.getUploadJobStatus(ctx, uploadResp.JobID)
		},
		PollInterval: opts.PollInterval,
		OnStatus:     opts.OnStatus,
	})
}

func contentTypeFromExt(ext string) string {
	switch ext {
	case ".geojson":
		return "application/geo+json"
	case ".json":
		return "application/json"
	case ".kml":
		return "application/vnd.google-earth.kml+xml"
	case ".gpx":
		return "application/gpx+xml"
	default:
		return ""
	}
}

// getPresignedURL requests a presigned URL for file upload.
func (c *Client) getPresignedURL(ctx context.Context, req PresignedURLRequest) (*PresignedURLResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.config.BaseURL+"/api/v1/storage/presigned-url",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	var result PresignedURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// startGeofenceUpload starts a geofence upload job.
func (c *Client) startGeofenceUpload(ctx context.Context, req UploadGeofencesRequest) (*UploadGeofencesResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.config.BaseURL+"/api/v1/geofences/upload",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	var result UploadGeofencesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result, nil
}

// getUploadJobStatus gets the status of a geofence upload job.
func (c *Client) getUploadJobStatus(ctx context.Context, jobID string) (*JobStatus, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		c.config.BaseURL+"/api/v1/geofences/upload/"+jobID+"/status",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	// Read body to parse into both raw map and structured response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse into structured fields
	var rawResp struct {
		JobID         string                 `json:"job_id"`
		Status        string                 `json:"status"`
		TotalFeatures int                    `json:"total_features"`
		CreatedCount  int                    `json:"created_count"`
		FailedCount   int                    `json:"failed_count"`
		ErrorMessage  *string                `json:"error_message"`
		Results       map[string]interface{} `json:"results"`
	}
	if err := json.Unmarshal(bodyBytes, &rawResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &JobStatus{
		JobID:         rawResp.JobID,
		Status:        rawResp.Status,
		TotalFeatures: rawResp.TotalFeatures,
		CreatedCount:  rawResp.CreatedCount,
		FailedCount:   rawResp.FailedCount,
		ErrorMessage:  rawResp.ErrorMessage,
		Results:       rawResp.Results,
	}, nil
}
