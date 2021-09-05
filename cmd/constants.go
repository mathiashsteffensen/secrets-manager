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
)

