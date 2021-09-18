package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"os"
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

func TestReadKeyFile(t *testing.T) {
	key := readKeyFile()

	assert.Equal(t, 32, len(key))
}
