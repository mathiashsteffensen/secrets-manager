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
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestDecrypt(t *testing.T) {
	key, err := GenRandomBytes(16)
	assert.Nil(t, err)

	plainContents := []byte("plain text contents")

	encryptedContents, err := Encrypt(plainContents, key)
	assert.Nil(t, err)

	// Test it can actually decrypt when used with the right key
	decryptedContents, err := Decrypt(encryptedContents, key)
	assert.Nil(t, err)
	assert.Equal(t, plainContents, decryptedContents)

	otherKey, err := GenRandomBytes(16)
	assert.Nil(t, err)

	// Test that it can't decrypt with another key
	nilSlice, err := Decrypt(encryptedContents, otherKey)
	assert.Nil(t, nilSlice)
	assert.Error(t, err)
}

func TestEncrypt(t *testing.T) {
	key, err := GenRandomBytes(16)
	assert.Nil(t, err)

	plainContents := []byte("plain text contents")

	encryptedContents, err := Encrypt(plainContents, key)
	assert.Nil(t, err)

	// Test that we can encrypt stuff
	assert.NotEqual(t, plainContents, encryptedContents)

	// Test that encrypted contents aren't the same with the same key (side effect of GCM block cipher)
	contentsEncryptedAgain, err := Encrypt(plainContents, key)
	assert.Nil(t, err)
	assert.NotEqual(t, encryptedContents, contentsEncryptedAgain)
}

func TestGenRandomBytes(t *testing.T) {
	byteLength := 32

	original, err := GenRandomBytes(byteLength)
	assert.Nil(t, err)

	assert.Equal(t, byteLength, len(original))

	for i := 0; i < 200_000; i++ {
		current, err := GenRandomBytes(byteLength)
		assert.Nil(t, err)
		assert.NotEqualf(t, string(original), string(current), "GenRandomBytes should not return the same thing again in a 200_000 iteration loop, failed on %s", strconv.Itoa(i))
	}
}

// Benchmarks
func BenchmarkEncrypt(b *testing.B) {
	key, _ := GenRandomBytes(16)

	plainText := []byte("plain text contents")
	for i := 0; i < b.N; i++ {

		Encrypt(plainText, key)
	}
}
