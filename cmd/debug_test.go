package cmd

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {

	// app run will return an error if the hostsfile does not exist
	t.Run("hostsfile that doesn't exist", func(t *testing.T) {
		args, _ := setup("-f", "no-existy", "debug")
		err := App.Run(args)
		assert.NotEmpty(t, err)
		if runtime.GOOS == "windows" {
			assert.Equal(t, "open no-existy: The system cannot find the file specified.", err.Error())
		} else {
			assert.Equal(t, "open no-existy: no such file or directory", err.Error())
		}
	})

	tests := map[string]struct {
		input    string
		expected string
	}{
		"hostsfile has only 1 line": {
			input: `127.0.0.1 localhost
`,
			expected: `hosts file path: test-debug
+-----------+-------+
|   TYPE    | COUNT |
+-----------+-------+
| lines     |     1 |
| entries   |     1 |
| comments  |     0 |
| empty     |     0 |
| malformed |     0 |
+-----------+-------+
`,
		},
		"hotsfile only has 1 malformed line": {
			input: `127.x.x.x localhost
`,
			expected: `hosts file path: test-debug
+-----------+-------+
|   TYPE    | COUNT |
+-----------+-------+
| lines     |     1 |
| entries   |     1 |
| comments  |     0 |
| empty     |     0 |
| malformed |     1 |
+-----------+-------+
`,
		},
		"hotsfile only has 1 empty line": {
			input: `
`,
			expected: `hosts file path: test-debug
+-----------+-------+
|   TYPE    | COUNT |
+-----------+-------+
| lines     |     1 |
| entries   |     0 |
| comments  |     0 |
| empty     |     1 |
| malformed |     0 |
+-----------+-------+
`,
		},
		"hotsfile only has 1 comment line": {
			input: `# comment
`,
			expected: `hosts file path: test-debug
+-----------+-------+
|   TYPE    | COUNT |
+-----------+-------+
| lines     |     1 |
| entries   |     0 |
| comments  |     1 |
| empty     |     0 |
| malformed |     0 |
+-----------+-------+
`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			defer read(t, "test-debug", test.input)()
			args, out := setup("-f", "test-debug", "debug")
			assert.Nil(t, App.Run(args))
			assert.Equal(t, test.expected, out.String())
		})
	}
}
