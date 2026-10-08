package tests

import (
	"os"
	"regexp"
	"testing"

	"github.com/spatialflow-io/spatialflow-go/v2/spatialflow"
)

func TestUserAgentReportsSDKVersion(t *testing.T) {
	if want := "spatialflow-go/" + spatialflow.SDKVersion; spatialflow.UserAgent != want {
		t.Errorf("UserAgent = %q, want %q", spatialflow.UserAgent, want)
	}
}

// The newest released heading is the first "## [x.y.z]" entry, which skips [Unreleased].
func TestSDKVersionMatchesNewestChangelogRelease(t *testing.T) {
	data, err := os.ReadFile("../CHANGELOG.md")
	if err != nil {
		t.Fatalf("read changelog: %v", err)
	}
	m := regexp.MustCompile(`(?m)^## \[(\d+\.\d+\.\d+)\]`).FindSubmatch(data)
	if m == nil {
		t.Fatal("no released version heading found in CHANGELOG.md")
	}
	if got := string(m[1]); got != spatialflow.SDKVersion {
		t.Errorf("SDKVersion = %q, newest CHANGELOG.md release = %q", spatialflow.SDKVersion, got)
	}
}
