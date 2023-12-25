package cmd

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	args, _ := setup("-f", "no-existy", "list")
	err := App.Run(args)
	assert.NotEmpty(t, err)
	if runtime.GOOS == "windows" {
		assert.Equal(t, "open no-existy: The system cannot find the file specified.", err.Error())
	} else {
		assert.Equal(t, "open no-existy: no such file or directory", err.Error())
	}

	file := "list"
	content := "127.0.0.1 localhost"

	cleanup := read(t, file, content)
	defer cleanup()

	args, out := setup("-f", file, "list")
	assert.Empty(t, App.Run(args))
	assert.Equal(t, content+"\n", out.String())
}
