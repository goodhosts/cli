package cmd

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	// app run will return an error if the hostsfile does not exist
	t.Run("hostsfile that doesn't exist", func(t *testing.T) {
		args, _ := setup("-f", "no-existy", "remove", "127.0.0.1")
		err := App.Run(args)
		assert.NotEmpty(t, err)
		if runtime.GOOS == "windows" {
			assert.Equal(t, "open no-existy: The system cannot find the file specified.", err.Error())
		} else {
			assert.Equal(t, "open no-existy: no such file or directory", err.Error())
		}
	})

	t.Run("no args", func(t *testing.T) {
		args, _ := setup("remove")
		assert.Equal(t, App.Run(args).Error(), "no input")
	})

	t.Run("remove an ip", func(t *testing.T) {
		defer write(t, "test-list", "127.0.0.1 localhost")()
		args, out := setup("-f", "test-list", "remove", "127.0.0.1")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "removing ip 127.0.0.1\n", out.String())

		args, out = setup("-f", "test-list", "list")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "", out.String())
	})

	t.Run("remove a host", func(t *testing.T) {
		defer write(t, "test-list", "127.0.0.1 localhost")()
		args, out := setup("-f", "test-list", "remove", "localhost")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "removing hostname localhost\n", out.String())

		args, out = setup("-f", "test-list", "list")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "", out.String())
	})

	t.Run("remove two hosts", func(t *testing.T) {
		defer write(t, "test-list", "127.0.0.1 localhost host1 host2")()
		args, out := setup("-f", "test-list", "remove", "127.0.0.1", "localhost", "host1")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "entry removed: localhost host1\n", out.String())

		args, out = setup("-f", "test-list", "list")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "127.0.0.1 host2\n", out.String())
	})

	t.Run("remove two hosts from multiple lines", func(t *testing.T) {
		defer write(t, "test-list", "127.0.0.1 localhost host1\n127.0.0.2 localhost host2")()
		args, out := setup("-f", "test-list", "remove", "host1", "host2")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "entry removed: host1 host2\n", out.String())

		args, out = setup("-f", "test-list", "list")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "127.0.0.1 localhost\n127.0.0.2 localhost\n", out.String())
	})
}
