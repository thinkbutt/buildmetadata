package buildmetadata

import (
	"testing"
)

func TestGitCommitSHA(t *testing.T) {
	sha := GitCommitSHA()
	if sha == "" {
		t.Fatalf("expected non-empty sha, got: %#v", sha)
	}
}
