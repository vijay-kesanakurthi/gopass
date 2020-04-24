package tests

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows.")
	}
	ts := newTester(t)
	defer ts.teardown()

	_, err := ts.run("move")
	assert.Error(t, err)

	ts.initStore()

	out, err := ts.run("move")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" mv old-path new-path\n", out)

	out, err = ts.run("move foo")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Usage: "+filepath.Base(ts.Binary)+" mv old-path new-path\n", out)

	out, err = ts.run("move foo bar")
	assert.Error(t, err)
	assert.Equal(t, "\nError: Source foo does not exist in source store : Entry is not in the password store\n", out)

	ts.initSecrets("")

	_, err = ts.run("move foo bar")
	assert.NoError(t, err)

	out, _ = ts.run("move foo/bar foo/baz")
	assert.Equal(t, "\nError: Source foo/bar does not exist in source store : Entry is not in the password store\n", out)

	_, err = ts.run("show -f bar/foo/bar")
	assert.NoError(t, err)

	_, err = ts.run("show -f baz")
	assert.NoError(t, err)
}
