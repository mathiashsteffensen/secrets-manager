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
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
// This simply prints a description and some instructions
var rootCmd = &cobra.Command{
	Use:   "secrets-manager",
	Short: "Manage application secrets",
	Long:  rootDescription,
}

var (
	secretsFile string
	keyFile     string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&secretsFile, "secrets-file", "s", "./config/secrets.yml.enc", "Secrets file to decrypt, edit and encrypt")
	rootCmd.PersistentFlags().StringVarP(&keyFile, "key-file", "k", "./config/master.key", "Encryption key file location")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
