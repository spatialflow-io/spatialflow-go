package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/spatialflow-io/spatialflow-go/v2/spatialflow"
)

func TestUploadGeofencesCompletesUploadBeforeImport(t *testing.T) {
	content := []byte(`{"type":"FeatureCollection","features":[]}`)
	file, err := os.CreateTemp(t.TempDir(), "upload-*.geojson")
	if err != nil {
		t.Fatalf("CreateTemp() error = %v", err)
	}
	if _, err := file.Write(content); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	var callOrder []string
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/storage/presigned-url":
			callOrder = append(callOrder, "presign")
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(
				w,
				`{"upload_url":%q,"file_id":"file-123","content_type":"application/geo+json"}`,
				server.URL+"/s3/upload",
			)
		case "/s3/upload":
			callOrder = append(callOrder, "s3_put")
			if r.Method != http.MethodPut {
				t.Errorf("S3 method = %s, want PUT", r.Method)
			}
			if r.ContentLength != int64(len(content)) {
				t.Errorf("S3 Content-Length = %d, want %d", r.ContentLength, len(content))
			}
			if got := r.Header.Get("Content-Type"); got != "application/geo+json" {
				t.Errorf("S3 Content-Type = %q, want application/geo+json", got)
			}
			if got := r.Header.Get("If-None-Match"); got != "*" {
				t.Errorf("S3 If-None-Match = %q, want *", got)
			}
			body, readErr := io.ReadAll(r.Body)
			if readErr != nil {
				t.Errorf("reading S3 body: %v", readErr)
			}
			if !reflect.DeepEqual(body, content) {
				t.Errorf("S3 body = %q, want %q", body, content)
			}
			w.WriteHeader(http.StatusOK)
		case "/api/v1/storage/uploads/file-123/complete":
			callOrder = append(callOrder, "complete")
			if r.Method != http.MethodPost {
				t.Errorf("complete method = %s, want POST", r.Method)
			}
			if got := r.Header.Get("Authorization"); got != "Bearer test-token" {
				t.Errorf("complete Authorization = %q, want Bearer test-token", got)
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"file_id":"file-123","status":"complete"}`)
		case "/api/v1/geofences/upload":
			callOrder = append(callOrder, "import")
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"job_id":"job-123","status":"pending"}`)
		case "/api/v1/geofences/upload/job-123/status":
			callOrder = append(callOrder, "status")
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"job_id":         "job-123",
				"status":         "completed",
				"created_count":  1,
				"failed_count":   0,
				"total_features": 1,
				"results":        map[string]interface{}{},
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client, err := spatialflow.NewClient(
		spatialflow.WithAccessToken("test-token"),
		spatialflow.WithBaseURL(server.URL),
	)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	result, err := client.UploadGeofences(context.Background(), spatialflow.UploadGeofencesOptions{
		FilePath:     file.Name(),
		Timeout:      5 * time.Second,
		PollInterval: time.Millisecond,
	})
	if err != nil {
		t.Fatalf("UploadGeofences() error = %v", err)
	}
	if result.Status != "completed" {
		t.Errorf("result status = %q, want completed", result.Status)
	}

	wantOrder := []string{"presign", "s3_put", "complete", "import", "status"}
	if !reflect.DeepEqual(callOrder, wantOrder) {
		t.Errorf("call order = %v, want %v", callOrder, wantOrder)
	}
}
