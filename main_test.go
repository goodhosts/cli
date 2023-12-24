package main

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	out := &bytes.Buffer{}
	logrus.SetOutput(out)

	args := os.Args[0:1]
	args = append(args, "version")
	assert.Nil(t, run(args))
	assert.Equal(t, "goodhosts dev@none built on unknown", out.String())
}

func TestDebug(t *testing.T) {
	out := &bytes.Buffer{}
	logrus.SetOutput(out)

	args := os.Args[0:1]
	args = append(args, "--debug", "version")
	assert.Nil(t, run(args))
	assert.True(t, strings.Contains(out.String(), "level=info msg=\"goodhosts dev@none built on unknown\""))
}

func TestQuiet(t *testing.T) {
	out := &bytes.Buffer{}
	logrus.SetOutput(out)

	args := os.Args[0:1]
	args = append(args, "--quiet", "version")
	assert.Nil(t, run(args))
	assert.Empty(t, out.String())
}
