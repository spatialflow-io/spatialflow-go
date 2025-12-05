package spatialflow

import (
	"context"
	"fmt"
	"time"
)

// JobResult represents the result of a completed job.
type JobResult struct {
	JobID            string
	Status           string
	CreatedCount     int
	FailedCount      int
	TotalFeatures    int
	Duration         time.Duration
	CreatedGeofences []CreatedGeofence
	Errors           []JobError
	Warnings         []string
	RawResponse      interface{}
}

// CreatedGeofence represents a successfully created geofence.
type CreatedGeofence struct {
	ID   string
	Name string
}

// JobError represents an error during job processing.
type JobError struct {
	Index int
	Name  string
	Error string
}

// JobStatus represents the status of an async job.
type JobStatus struct {
	JobID         string
	Status        string
	TotalFeatures int
	CreatedCount  int
	FailedCount   int
	ErrorMessage  *string
	Results       map[string]interface{}
}

// PollJobOptions configures job polling.
type PollJobOptions struct {
	// JobID is the job identifier.
	JobID string

	// FetchStatus is the function to fetch job status.
	FetchStatus func(ctx context.Context) (*JobStatus, error)

	// PollInterval is the interval between status checks (default 2s).
	PollInterval time.Duration

	// TerminalStates are the states that indicate job completion.
	TerminalStates []string

	// OnStatus is called after each status check.
	OnStatus func(status string, response *JobStatus)
}

// PollJob polls a job until completion or context cancellation.
func (c *Client) PollJob(ctx context.Context, opts PollJobOptions) (*JobResult, error) {
	if opts.PollInterval == 0 {
		opts.PollInterval = 2 * time.Second
	}
	if len(opts.TerminalStates) == 0 {
		opts.TerminalStates = []string{"completed", "failed"}
	}

	startTime := time.Now()
	ticker := time.NewTicker(opts.PollInterval)
	defer ticker.Stop()

	// Initial fetch
	response, err := opts.FetchStatus(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch job status: %w", err)
	}

	for {
		status := response.Status
		jobID := response.JobID

		if opts.OnStatus != nil {
			opts.OnStatus(status, response)
		}

		if isTerminal(status, opts.TerminalStates) {
			if status == "failed" {
				errMsg := "unknown error"
				if response.ErrorMessage != nil {
					errMsg = *response.ErrorMessage
				}
				return nil, &JobFailedError{
					JobID:   jobID,
					Message: errMsg,
				}
			}
			return buildJobResult(response, time.Since(startTime)), nil
		}

		// Wait for next poll or context cancellation
		select {
		case <-ctx.Done():
			return nil, &JobTimeoutError{
				JobID:      jobID,
				LastStatus: status,
				Err:        ctx.Err(),
			}
		case <-ticker.C:
			response, err = opts.FetchStatus(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to fetch job status: %w", err)
			}
		}
	}
}

func isTerminal(status string, terminals []string) bool {
	for _, t := range terminals {
		if status == t {
			return true
		}
	}
	return false
}

func buildJobResult(resp *JobStatus, duration time.Duration) *JobResult {
	result := &JobResult{
		JobID:         resp.JobID,
		Status:        resp.Status,
		CreatedCount:  resp.CreatedCount,
		FailedCount:   resp.FailedCount,
		TotalFeatures: resp.TotalFeatures,
		Duration:      duration,
		RawResponse:   resp,
	}

	// Extract results if available (with safe type assertions)
	if resp.Results != nil {
		if created, ok := resp.Results["created_geofences"].([]interface{}); ok {
			for _, c := range created {
				if m, ok := c.(map[string]interface{}); ok {
					result.CreatedGeofences = append(result.CreatedGeofences, CreatedGeofence{
						ID:   safeString(m["id"]),
						Name: safeString(m["name"]),
					})
				}
			}
		}
		if errs, ok := resp.Results["errors"].([]interface{}); ok {
			for _, e := range errs {
				if m, ok := e.(map[string]interface{}); ok {
					result.Errors = append(result.Errors, JobError{
						Index: safeInt(m["index"]),
						Name:  safeString(m["name"]),
						Error: safeString(m["error"]),
					})
				}
			}
		}
		if warns, ok := resp.Results["warnings"].([]interface{}); ok {
			for _, w := range warns {
				result.Warnings = append(result.Warnings, fmt.Sprint(w))
			}
		}
	}

	return result
}

// safeString extracts a string from interface{}, returns "" if not a string.
func safeString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	if v != nil {
		return fmt.Sprint(v)
	}
	return ""
}

// safeInt extracts an int from interface{}, handles float64 (JSON default) and int.
func safeInt(v interface{}) int {
	switch n := v.(type) {
	case float64:
		return int(n)
	case int:
		return n
	case int64:
		return int(n)
	default:
		return 0
	}
}

// JobTimeoutError indicates a job timed out or context was cancelled.
type JobTimeoutError struct {
	JobID      string
	LastStatus string
	Err        error
}

func (e *JobTimeoutError) Error() string {
	return fmt.Sprintf("job %s: %v (last status: %s)", e.JobID, e.Err, e.LastStatus)
}

func (e *JobTimeoutError) Unwrap() error {
	return e.Err
}

// JobFailedError indicates a job failed.
type JobFailedError struct {
	JobID   string
	Message string
}

func (e *JobFailedError) Error() string {
	return fmt.Sprintf("job %s failed: %s", e.JobID, e.Message)
}
