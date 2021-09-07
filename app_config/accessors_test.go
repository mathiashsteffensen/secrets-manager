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

package AppConfig

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")

func TestGet(t *testing.T) {
	got, err := Get("secret")
	assert.Nil(t, err)
	assert.Equal(t, "hello", got)

	got, err = Get("super.deeply.nested")
	assert.Nil(t, err)
	assert.Equal(t, got, "value")

	_, err = Get("this.is.not.a.real.key")
	assert.Error(t, err)
}

func TestGetOrDefault(t *testing.T) {
	got := GetOrDefault("secret", "default")
	assert.Equal(t, "hello", got)

	got = GetOrDefault("super.deeply.nested", "default")
	assert.Equal(t, got, "value")

	got = GetOrDefault("this.is.not.a.real.key", "default")
	assert.Equal(t, got, "default")
}

func TestMustGet(t *testing.T) {
	got := MustGet("secret")
	assert.Equal(t, "hello", got)

	got = MustGet("super.deeply.nested")
	assert.Equal(t, got, "value")

	defer func() {
		r := recover()

		assert.NotNil(t, r)
	}()
	_ = MustGet("this.is.not.a.real.key")
}

func TestGetConfig(t *testing.T) {
	expectedConfig := Config{
		"secret": "hello",
		"super": Config{
			"deeply": Config{
				"nested": "value",
			},
		},
	}

	actualConfig := GetConfig()

	assert.Equal(t, expectedConfig, actualConfig)
}

func TestExists(t *testing.T) {
	got := Exists("secret")
	assert.Equal(t, true, got)

	got = Exists("this.is.not.a.real.key")
	assert.Equal(t, false, got)
}

func TestAllKeys(t *testing.T) {
	for _, s := range AllKeys() {
		assert.Equal(t, true, Exists(s))
	}
}
