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

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"github.com/lucsky/cuid"
	"io"
	"io/ioutil"
	mathRand "math/rand"
	"path/filepath"
)

func init() {
	seed := binary.BigEndian.Uint64([]byte(cuid.New()))
	mathRand.Seed(int64(seed))
}

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

func GenRandomBytes(byteLength int) (randomBytes []byte, err error) {
	randomBytes = make([]byte, byteLength)

	_, err = mathRand.Read(randomBytes)

	return
}
