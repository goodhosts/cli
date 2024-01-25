package cmd

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// app run will return an error if the hostsfile does not exist
	t.Run("hostsfile that doesn't exist", func(t *testing.T) {
		args, _ := setup("-f", "no-existy", "list")
		err := App.Run(args)
		assert.NotEmpty(t, err)
		if runtime.GOOS == "windows" {
			assert.Equal(t, "open no-existy: The system cannot find the file specified.", err.Error())
		} else {
			assert.Equal(t, "open no-existy: no such file or directory", err.Error())
		}
	})

	// this is really a noop but future proofs us a bit
	t.Run("default action", func(t *testing.T) {
		defer read(t, "test-list", "127.0.0.1 localhost")()
		args, out := setup("-f", "test-list")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "127.0.0.1 localhost"+"\n", out.String())
	})

	// this is really a noop but future proofs us a bit
	t.Run("hostsfile that is write accessible", func(t *testing.T) {
		defer write(t, "test-list", "127.0.0.1 localhost")()
		args, out := setup("-f", "test-list", "list")
		assert.Empty(t, App.Run(args))
		assert.Equal(t, "127.0.0.1 localhost"+"\n", out.String())
	})

	// test reading the system hostsfile
	t.Run("read the system hostsfile", func(t *testing.T) {
		args, out := setup("list")
		assert.Empty(t, App.Run(args))
		assert.NotEmpty(t, out)
	})

	tests := map[string]struct {
		input    string
		expected string
	}{
		"hostsfile that is readonly": {
			input: `127.0.0.1 localhost
`,
			expected: `127.0.0.1 localhost
`,
		},
		"malformed comment added to row": {
			input: `127.x.x.x localhost
127.x.x.x localhost # comment
`,
			expected: `127.x.x.x localhost # <<< Malformed!
127.x.x.x localhost # comment <<< Malformed!
`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer read(t, "list-test", test.input)()
			args, out := setup("-f", "list-test", "list")
			assert.Empty(t, App.Run(args))
			assert.Equal(t, test.expected, out.String())
		})
	}

}
