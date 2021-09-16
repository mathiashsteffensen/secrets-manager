package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestRunGKeyCmd(t *testing.T) {
	validKeyLengths := []string{
		"16",
		"24",
		"32",
	}

	keyFile = "./master.key"

	cobraCmd := cobra.Command{}

	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()

	for _, keyLength := range validKeyLengths {
		args := []string{keyLength}

		runGKeyCmd(&cobraCmd, args)

		key, err := ioutil.ReadFile(keyFile)
		assert.Nil(t, err)

		assert.Equal(t, keyLength, strconv.Itoa(len(key)))

		err = os.Remove(keyFile)
		assert.Nil(t, err)
	}

	keyFile = "../config/master.key"
}
