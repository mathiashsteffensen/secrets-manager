package FileHelpers

import (
	"encoding/base64"
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"io/ioutil"
	"path/filepath"
)

func ReadEncryptedSecretsFile(fileLocation string, key []byte) (decrypted []byte, err error) {
	secrets, err := LoadFile(fileLocation)
	if err != nil {
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(string(secrets))
	if err != nil {
		return
	}

	decrypted, err = crypto.Decrypt(decoded, key)

	return
}

func LoadFile(relativePath string) (contents []byte, err error) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return
	}

	contents, err = ioutil.ReadFile(absolutePath)
	return
}
