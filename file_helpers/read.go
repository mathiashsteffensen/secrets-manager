package FileHelpers

import (
	"encoding/base64"
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"io/ioutil"
	"path/filepath"
)

func ReadEncryptedSecretsFile(fileLocation string, key []byte) (decrypted []byte, err error) {
	absSecretsFile, err := filepath.Abs(fileLocation)
	if err != nil {
		return
	}

	secrets, err := ioutil.ReadFile(absSecretsFile)
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
