package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
)

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
