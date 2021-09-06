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

const (
	rootDescription = `
  CLI tool for editing and encrypting secrets files.

  Checkout https://github.com/mathiashsteffensen/secrets-manager for detailed instructions
`
	editDescription = `
  Edit your applications secrets file, by default looks in ./config/secrets.yml.enc from the directory where it's called

  Usage with custom config and key file:
    secrets-manager edit -s ./config/credentials.yml.enc -k ./config/secrets.key
`
	gkeyDescription = `
  Generates a new master encryption key, by default is 32 bytes to use AES-256 encryption.
  The key is saved to ./config/master.key by default
	
  To generate a 24 byte encryption key:
    secrets-manager g:key 24
`
	printDescription = `
  Prints all config values formatted as json

  Example usage:
    secrets-manager load ./config/master.key -s ./config/secrets.yml.enc
`
)
