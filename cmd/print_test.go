package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunPrintCmd(t *testing.T) {
	cmd := &cobra.Command{}
	var args []string

	defer func() {
		r := recover()
		assert.Nil(t, r)
	}()

	runPrintCmd(cmd, args)
}
