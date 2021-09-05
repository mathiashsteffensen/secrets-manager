package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
)

func DecryptSecrets(secrets []byte, key []byte) (decrypted []byte, err error) {
	gcm, err := NewGCM(key)

	nonce := secrets[:gcm.NonceSize()]
	decrypted, err = gcm.Open(nil, nonce, secrets[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func NewGCM(key []byte) (gcm cipher.AEAD, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err = cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return
}

type Encryptor func([]byte, string) error

func NewEncryptor(key []byte) Encryptor {
	return func(contents []byte, location string) (err error) {
		gcm, err := NewGCM(key)
		if err != nil {
			return
		}

		nonce := make([]byte, gcm.NonceSize())
		_, err = io.ReadFull(rand.Reader, nonce)
		if err != nil {
			return
		}

		encrypted :=gcm.Seal(nonce, nonce, contents, nil)

		absSecretsFile, err := filepath.Abs(location)
		if err != nil {
			return
		}

		err = ioutil.WriteFile(absSecretsFile, encrypted, 0777)
		if err != nil {
			return
		}

		fmt.Printf("Saved encrypted secrets to %s\n", absSecretsFile)

		return
	}
}
