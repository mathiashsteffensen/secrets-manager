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
	"fmt"
	AppConfig "github.com/mathiashsteffensen/secrets-manager/app_config"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a key from your configuration and print it to stdout",
	Run:   runGetCmd,
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringSliceVarP(&envFiles, "env-files", "f", []string{}, "Unencrypted environment yml files to load")
}

func runGetCmd(_ *cobra.Command, args []string) {
	if len(args) != 1 {
		logger.Printf("Please provide a single key to get\n")
		return
	}

	if secretsFile != "" && keyFile != "" {
		err := AppConfig.LoadEncrypted(secretsFile, keyFile)
		cobra.CheckErr(err)
	}

	for _, file := range envFiles {
		err := AppConfig.Load(file)
		cobra.CheckErr(err)
	}

	value := AppConfig.MustGet(args[0])

	fmt.Print(value)
}
