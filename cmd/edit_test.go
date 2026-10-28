package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestRunEditCmd(t *testing.T) {
	cobraCmd := cobra.Command{}
	var args []string

	err := os.Setenv("EDITOR", "cat")
	assert.Nil(t, err)

	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()

	runEditCmd(&cobraCmd, args)

	secretsFile = "../config/config.yml.enc"

	runEditCmd(&cobraCmd, args)

	err = os.Remove(secretsFile)
	assert.Nil(t, err)

	secretsFile = "../config/secrets.yml.enc"
}

func TestOpenTempFileConnectsEditorToTerminal(t *testing.T) {
	dir := t.TempDir()

	editorDir := filepath.Join(dir, "my editor")
	err := os.Mkdir(editorDir, 0755)
	assert.Nil(t, err)

	// Stands in for a terminal editor: it writes what the user types into the file
	editor := filepath.Join(editorDir, "editor.sh")
	err = os.WriteFile(editor, []byte("#!/bin/sh\ncat > \"$1\"\n"), 0755)
	assert.Nil(t, err)

	typed := filepath.Join(dir, "typed")
	err = os.WriteFile(typed, []byte("key: value\n"), 0644)
	assert.Nil(t, err)

	stdin, err := os.Open(typed)
	assert.Nil(t, err)
	defer stdin.Close()

	originalStdin := os.Stdin
	os.Stdin = stdin
	defer func() { os.Stdin = originalStdin }()

	// EDITOR is parsed by the shell: quoted paths with spaces and padding both work
	t.Setenv("EDITOR", `  "`+editor+`"  `)

	// The file path must reach the editor as a single argument
	content := openTempFile(filepath.Join(dir, "secrets edit.yml"))

	assert.Equal(t, "key: value\n", string(content))
}

func TestReadKeyFile(t *testing.T) {
	key := readKeyFile()

	assert.Equal(t, 32, len(key))
}
