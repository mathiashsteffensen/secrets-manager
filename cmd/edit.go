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
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	secretsFile string
	keyFile     string
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit your application secrets file",
	Long:  editDescription,
	Run:   runEditCmd,
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&secretsFile, "secrets-file", "s", "./config/secrets.yml.enc", "Secrets file to decrypt, edit and encrypt")
	editCmd.Flags().StringVarP(&keyFile, "key-file", "k", "./config/master.key", "Encryption key file location")
}

func runEditCmd(cmd *cobra.Command, args []string) {
	key := readKeyFile()

	secrets := readEncryptedSecretsFile()

	if secrets != nil {
		decryptedSecrets, err := crypto.Decrypt(secrets, key)
		cobra.CheckErr(err)
		secrets = decryptedSecrets
	}

	dir, err := ioutil.TempDir(".", "tmp")
	cobra.CheckErr(err)

	defer os.RemoveAll(dir) // clean up

	tmpFile := createTempFile(secrets, dir)

	plainTextContent := openTempFile(tmpFile)

	encryptedContent, err := crypto.Encrypt(plainTextContent, key)
	cobra.CheckErr(err)

	location := saveEncryptedSecretsFile(encryptedContent)

	logger.Printf("Saved encrypted secrets to %s\n", location)
}

func readKeyFile() []byte {
	absKeyFile, err := filepath.Abs(keyFile)
	cobra.CheckErr(err)

	key, err := ioutil.ReadFile(absKeyFile)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			logger.Printf("\nNo key file exists in the specified location %s\n\n", absKeyFile)
			logger.Println("To create new key file run: secrets-manager g:key")
			cobra.CheckErr(err)
		} else {
			cobra.CheckErr(err)
		}

		return nil
	}

	return key
}

func readEncryptedSecretsFile() []byte {
	absSecretsFile, err := filepath.Abs(secretsFile)
	cobra.CheckErr(err)

	secrets, err := ioutil.ReadFile(absSecretsFile)

	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil
		} else {
			cobra.CheckErr(err)
		}
	}

	return secrets
}

func saveEncryptedSecretsFile(content []byte) (location string) {
	location, err := filepath.Abs(secretsFile)
	cobra.CheckErr(err)

	err = ioutil.WriteFile(location, content, 0777)
	cobra.CheckErr(err)

	return
}

func createTempFile(content []byte, dir string) string {
	tmp := filepath.Join(dir, "secrets.edit.yml")

	if err := ioutil.WriteFile(tmp, content, 0666); err != nil {
		logger.Fatal(err)
	}

	return tmp
}

func openTempFile(location string) []byte {
	editor := env("EDITOR", "subl -w")
	editorSlice := strings.Split(editor, " ")

	cmd := exec.Command(editorSlice[0], editorSlice[1], location)

	err := cmd.Start()
	cobra.CheckErr(err)

	err = cmd.Wait()
	cobra.CheckErr(err)

	content, err := ioutil.ReadFile(location)
	cobra.CheckErr(err)

	return content
}
