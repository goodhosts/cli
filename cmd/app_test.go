package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func TestDebug(t *testing.T) {
	args, out := setup("--debug", "version")
	Version("test-version", "test-commit", "test-date")
	assert.Nil(t, App.Run(args))
	assert.True(t, strings.Contains(out.String(), "level=info msg=\"goodhosts test-version@test-commit built on test-date\""))

	// reset logrus for other tests
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&easy.Formatter{
		LogFormat: "%msg%",
	})
}

func TestQuiet(t *testing.T) {
	args, out := setup("--quiet", "version")
	assert.Nil(t, App.Run(args))
	assert.Empty(t, out.String())
}

// Setup takes args slice and resets logrus and returns args to pass to App.Run
func setup(args ...string) ([]string, *bytes.Buffer) {
	out := &bytes.Buffer{}
	logrus.SetOutput(out)
	a := os.Args[0:1]
	a = append(a, args...)
	return a, out
}

func read(t *testing.T, name, file string) func() {
	err := os.WriteFile(name, []byte(file), 0440)
	assert.Nil(t, err)
	return func() {
		assert.Nil(t, os.Remove(name))
	}
}

//func write(t *testing.T, name, file string) func() {
//	err := os.WriteFile(name, []byte(file), 0770)
//	assert.Nil(t, err)
//	return func() {
//		assert.Nil(t, os.Remove(name))
//	}
//}
