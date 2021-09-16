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

// Package crypto provides helpers for encrypting and decrypting byte slices using a pseudo-random seed base on CUIDs
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"github.com/lucsky/cuid"
	mathRand "math/rand"
)

func init() {
	// Seed with pseudo random CUID
	seed := binary.BigEndian.Uint64([]byte(cuid.New()))
	mathRand.Seed(int64(seed))
}

// Decrypt is a function that takes a byte slice of contents to decrypt and a key to use for the decryption
// it returns the decrypted result and any eventual error
func Decrypt(secrets []byte, key []byte) (decrypted []byte, err error) {
	gcm, err := NewGCM(key)
	if err != nil {
		return nil, err
	}

	nonce := secrets[:gcm.NonceSize()]
	decrypted, err = gcm.Open(nil, nonce, secrets[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// Encrypt is a function that takes a byte slice of contents to encrypt and a key to use for the encryption
// it returns the encrypted result and any eventual error
func Encrypt(contents []byte, key []byte) (encryptedContents []byte, err error) {
	gcm, err := NewGCM(key)
	if err != nil {
		return
	}

	nonce, err := GenRandomBytesBase64(gcm.NonceSize())
	if err != nil {
		return
	}

	encryptedContents = gcm.Seal(nonce, nonce, contents, nil)

	return
}

// NewGCM Takes an encryption/decryption key and returns a new GCM based un an underlying AES block cipher
// The GCM is used to perform cryptographic operations, it is what performs the actual encryption and decryption of our data
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

// GenRandomBytes generates a byte slice filled with a specified number of pseudo random bytes
func GenRandomBytes(sliceLength int) (randomBytes []byte, err error) {
	randomBytes = make([]byte, sliceLength)

	_, err = mathRand.Read(randomBytes)

	return
}

// GenRandomBytesBase64 generates a byte slice filled with a specified number of pseudo random bytes - base64 encoded
func GenRandomBytesBase64(sliceLength int) (randomBytes []byte, err error) {
	randomBytes, err = GenRandomBytes(sliceLength)
	if err != nil {
		return
	}

	randomBytes = []byte(base64.StdEncoding.EncodeToString(randomBytes))[:sliceLength]

	return
}

