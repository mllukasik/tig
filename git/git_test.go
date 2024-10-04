package git

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecutionExpectNoErrors(t *testing.T) {
	branch := parseBranchLine("* main                f5b4e88 [origin/main] create prune command")
	assert.Equal(t, "main", branch.Name)
	assert.Equal(t, true, branch.Current)
	assert.Equal(t, true, branch.RemoteTracking)
	assert.Equal(t, "[origin/main]", branch.RemoteName)
	assert.Equal(t, "f5b4e88", branch.LastCommitMessage)
	assert.Equal(t, "create prune command", branch.LastCommitHash)
}

func TestExecutionExpectNoErrors1(t *testing.T) {
	branch := parseBranchLine("  test                   f5b4e88 create prune command")
	assert.Equal(t, "test", branch.Name)
	assert.Equal(t, false, branch.Current)
	assert.Equal(t, false, branch.RemoteTracking)
	assert.Equal(t, "", branch.RemoteName)
	assert.Equal(t, "f5b4e88", branch.LastCommitMessage)
	assert.Equal(t, "create prune command", branch.LastCommitHash)
}
