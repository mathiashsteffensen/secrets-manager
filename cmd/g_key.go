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
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

// gKeyCmd represents the g:key command
var gKeyCmd = &cobra.Command{
	Use:   "g:key",
	Short: "Generate a new master key file",
	Long:  gkeyDescription,
	Run:   runGKeyCmd,
}

func init() {
	rootCmd.AddCommand(gKeyCmd)
}

func runGKeyCmd(cmd *cobra.Command, args []string) {
	var byteLength string

	if len(args) != 0 {
		byteLength = args[0]
	} else {
		byteLength = "32"
	}

	fmt.Printf("  Using byte length %s for master key\n", byteLength)

	intByteLength, err := strconv.Atoi(byteLength)

	cobra.CheckErr(err)

	if intByteLength != 32 && intByteLength != 24 && intByteLength != 16 {
		fmt.Println("  byte length should be one of [32, 24, 16]")
		return
	}

	absKeyFile, err := filepath.Abs(keyFile)

	cobra.CheckErr(err)

	randomBytes, err := crypto.GenRandomBytesBase64(intByteLength)

	cobra.CheckErr(err)

	err = ioutil.WriteFile(absKeyFile, randomBytes, 0777)

	cobra.CheckErr(err)

	fmt.Println("  Created new key file at:", keyFile)
	fmt.Println("  If using git for version control remember to add this file to your .gitignore")
}
