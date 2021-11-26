# Go Secrets Manager

[![CircleCI](https://circleci.com/gh/mathiashsteffensen/secrets-manager/tree/master.svg?style=shield)](https://circleci.com/gh/mathiashsteffensen/secrets-manager/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mathiashsteffensen/secrets-manager)](https://goreportcard.com/report/github.com/mathiashsteffensen/secrets-manager)

#### !IMPORTANT: The API is not stable, this project is still in development and shouldn't be used in production until its v1.0.0 release

#### !IMPORTANT: I am not a cryptography expert, just a programmer who likes experimenting. Most of the hardcore cryptography is handled by the Go standard library, for details on how things are encrypted see the crypto package.

This command line and associated Go API is meant to make managing secrets in your applications easier.

It provides a CLI tool for editing and encrypting a secrets.yml.enc file (or whatever file name you want to configure, more on that later). As well as a Go API for accessing this file in your application. The Go API also provides options to load other unencrypted yaml files to make all configuration accessible from one place.

Files are encrypted using AES (https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)

Main features:
* CLI tools for editing and managing encrypted secrets
* AppConfig
  * Load encrypted YAML files
  * Merge encrypted secrets with other YAML config files
  * Check if keys exist before attempting to fetch their values for easy error handling
  * Alternatively just call AppConfig.MustGet("key") and let it panic
  * Fetch deeply nested keys with simple dot notation AppConfig.MustGet("some_api.read_permission.token")

## The CLI Tool

Installation:
```bash
go install github.com/mathiashsteffensen/secrets-manager@latest
```

Building latest from source (requires just git, make* & Go):
```bash
git clone git@github.com:mathiashsteffensen/secrets-manager.git
cd secrets-manager
make install
make build # Binary is located in ./dist/secrets-manager
```
(* It doesn't really 'require' make, alternatively run 'go mod download && go build -o dist/secrets-manager')

Usage:
```bash
secrets-manager # Print available commands
EDITOR="subl -w" secrets-manager edit # Edit encrypted secrets file
secrets-manager g:key # Generate a new master key
```

### secrets-manager edit
This command opens your decrypted secrets file for editing in the editor saved in the EDITOR env variable, by default it will attempt to open this file in the sublime text editor, 
this default behavior is equivalent to setting EDITOR="subl -w" in your environment

Options and default values:

Description | Flag | Default value | Example
--- | --- | --- | ---
Secrets file to save encrypted contents too | -s, --secrets-file | ./config/secrets.yml.enc | `secrets-manager edit -s .config/credentials.yml.enc`
Key file to use for encryption and decryption | -k, --key-file | ./config/master.key | `secrets-manager edit -k .config/secret.key`

### secrets-manager g:key
Options and default values:

Description | Flag | Default value | Example
--- | --- | --- | ---
Byte length to use for encryption key, must be one of [16, 24, 32] | none (positional) | 32 | `secrets-manager g:key 24`
Key file to use for encryption and decryption | -k, --key-file | ./config/master.key | `secrets-manager g:key 24 -k .config/secret.key`

### secrets-manager print
This command loads encrypted and unencrypted configuration files and prints them in JSON format to STDOUT.
Useful for scripting and easy parsing.

Options and default values:

Description | Flag | Default value | Example
--- | --- | --- | ---
Secrets file to save encrypted contents too | -s, --secrets-file | ./config/secrets.yml.enc | `secrets-manager print -s .config/credentials.yml.enc`
Key file to use for encryption and decryption | -k, --key-file | ./config/master.key | `secrets-manager print -k .config/secret.key`
Other unencrypted environment files | -f, --env-files | [./config/env.yml] | `secrets-manager print -f [./config/env.production.yml, ./config/env.development.yml]`

## The Go API

Installation:
```bash
go get github.com/mathiashsteffensen/secrets-manager/app_config
```

Documentation: https://pkg.go.dev/github.com/mathiashsteffensen/secrets-manager/app_config

Usage:

```go
package main

import (
    "fmt"
    "github.com/mathiashsteffensen/secrets-manager/app_config"
)

func init() {
    // Load encrypted file using specified key file
	// When GO_ENV is set to production the key file location will be ignored and instead the GO_MASTER_KEY env variable will be used
    err := AppConfig.LoadEncrypted("./config/secrets.yml.enc", "./config/master.key")
    if err != nil {
        panic(err)
    }
    // Load other YAML files into AppConfig
    // This is useful to keep all your configuration values in one place, this will not override other values if the keys already exist
    err = AppConfig.Load("./config/env.yml")
    if err != nil {
        panic(err)
    }
}

func main() {
    // Check if a key exists before calling AppConfig.MustGet
    if AppConfig.Exists("secret") {
        secret := AppConfig.MustGet("secret")

		// Assuming a yaml file containing
		//  development:
		//      secret: Hello World
		// And GO_ENV set to the default of "development"
        fmt.Println(secret)
		// Prints "Hello World" to stdout
    }
    
    // Or call AppConfig.Get and handle error
    secret, err := AppConfig.Get("secret")
    
    if err != nil {
        panic(err)
    }

    fmt.Println(secret)
}
```
