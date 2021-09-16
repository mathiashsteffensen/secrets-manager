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
	"encoding/json"
	"fmt"
	AppConfig "github.com/mathiashsteffensen/secrets-manager/app_config"
	"github.com/spf13/cobra"
)

var (
	// printCmd represents the loadEnv command
	printCmd = &cobra.Command{
		Use:  "print",
		Long: printDescription,
		Run:  runPrintCmd,
	}

	envFiles []string
)

func init() {
	rootCmd.AddCommand(printCmd)

	printCmd.Flags().StringVarP(&secretsFile, "secrets-file", "s", "./config/secrets.yml.enc", "Secrets file to decrypt, edit and encrypt")
	printCmd.Flags().StringVarP(&keyFile, "key-file", "k", "./config/master.key", "Encryption key file location")
	printCmd.Flags().StringSliceVarP(&envFiles, "env-files", "f", []string{"./config/env.yml"}, "Unencrypted environment yml files to load")
}

func runPrintCmd(_ *cobra.Command, _ []string) {
	err := AppConfig.LoadEncrypted(secretsFile, keyFile)
	cobra.CheckErr(err)

	for _, file := range envFiles {
		err = AppConfig.Load(file)
		cobra.CheckErr(err)
	}

	jsonBytes, err := json.MarshalIndent(AppConfig.GetConfig(), "", "    ")
	cobra.CheckErr(err)

	fmt.Println(string(jsonBytes))
}
