# Go Secrets Manager

This command line and associated Go API is meant to make managing secrets in your applications easier.

It provides a CLI tool for editing and encrypting a secrets.yml.enc file (or whatever file name you want to configure, more on that later). As well as a Go API for accessing this file in your application. The Go API also provides options to load other unencrypted yaml files to make all configuration accessible from one place.

Files are encrypted using AES (https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)

Main features:
* CLI tools for editing and managing encrypted secrets
* AppConfig
  * Load AES encrypted files independent of the CLI tool, just provide an encrypted secrets file location and a key file location
  * Merge encrypted secrets with other YAML config files
  * Check if keys exist before attempting to fetch their values for easy error handling
  * Alternatively just call AppConfig.MustGet("key") and let it panic
  * Fetch deeply nested keys with simple dot notation AppConfig.MustGet("some_api.read_permission.token")
  * By default

## The CLI Tool

Installation:
```bash
go install github.com/mathiashsteffensen/secrets-manager@latest
```

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
Byte length to use for encryption key, must be one of [16, 24, 32] | none (positional) | ./config/secrets.yml.enc | `secrets-manager g:key 24`
Key file to use for encryption and decryption | -k, --key-file | ./config/master.key | `secrets-manager g:key 24 -k .config/secret.key`

## The Go API

Installation:
```bash
go get github.com/mathiashsteffensen/secrets-manager/app_config
```

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
    // Load other YAMsL files into AppConfig
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
