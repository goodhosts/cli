package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/goodhosts/cli/cmd"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// Run exported in tests to let subcommands call the main run
// Setup takes args slice and resets logrus and returns args to pass to App.Run
func setup(args ...string) ([]string, *bytes.Buffer) {
	out := &bytes.Buffer{}
	logrus.SetOutput(out)
	a := os.Args[0:1]
	a = append(a, args...)
	return a, out
}

func TestVersion(t *testing.T) {
	// test version passed in run()
	args, out := setup("version")
	assert.Nil(t, run(args))
	assert.Equal(t, "goodhosts dev@none built on unknown", out.String())

	// test that version@commit + date work
	args, out = setup("version")
	cmd.Version("test-version", "test-commit", "test-date")
	assert.Nil(t, cmd.App.Run(args))
	assert.Equal(t, "goodhosts test-version@test-commit built on test-date", out.String())

	// reset for future tests
	cmd.Version("dev", "none", "unknown")
}
