/*
Copyright © 2021 Mathias H Steffensen mathiashsteffensen@protonmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"errors"
	AppConfig "github.com/mathiashsteffensen/secrets-manager/app_config"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

var (
	// loadEnvCmd represents the loadEnv command
	loadEnvCmd = &cobra.Command{
		Use:   "load-env",
		Long: `
  Example usage:

  secrets-manager load-env -k ./config/master.key -s ./config/secrets.yml.enc go run main.go`,
		Run: runLoadEnvCmd,
	}

	envFiles []string
)

func init() {
	rootCmd.AddCommand(loadEnvCmd)

	loadEnvCmd.Flags().StringVarP(&secretsFile, "secrets-file", "s", "./config/secrets.yml.enc", "Secrets file to decrypt, edit and encrypt")
	loadEnvCmd.Flags().StringVarP(&keyFile, "key-file", "k", "./config/master.key", "Encryption key file location")
	loadEnvCmd.Flags().StringSliceVarP(&envFiles, "env-files", "f", []string{"./config/env.yml"}, "Unencrypted environment yml files to load")
}

func runLoadEnvCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cobra.CheckErr(errors.New("no command to run"))
	}

	err := AppConfig.LoadEncrypted(secretsFile, keyFile)
	cobra.CheckErr(err)

	err = AppConfig.Load("./config/env.yml")
	cobra.CheckErr(err)

	keys := AppConfig.AllKeys()

	for i, key := range keys {
		key = strings.ToUpper(key)
		key = strings.Replace(key, ".", "_", -1)
		value, err := AppConfig.Get(keys[i])
		cobra.CheckErr(err)

		err = os.Setenv(key, value.(string))
		cobra.CheckErr(err)

		keys[i] = key
	}

	var executable *exec.Cmd

	if len(args) == 1 {
		executable = exec.Command(args[0])
	} else {
		baseCommand := args[0]
		commandArgs := args[1:]

		executable = exec.Command(baseCommand, commandArgs...)
	}

	executable.Stdout = os.Stdout
	executable.Stderr = os.Stderr
	executable.Stdin = os.Stdin

	err = executable.Run()
	cobra.CheckErr(err)
}
