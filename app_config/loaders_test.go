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
	"os"
	"testing"
)

func TestLoadEncrypted(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")
	assert.Nil(t, err)

	devConfig := config["development"].(map[string]interface{})
	assert.Equal(t, "hello", devConfig["secret"])

	err = os.Setenv("GO_ENV", "production")
	assert.Nil(t, err)

	err = LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")
	assert.ErrorIs(t, err, ErrNoGoMasterKey)

	err = os.Setenv("GO_MASTER_KEY", "mvemS1oqdY3fUWNE/DVpkWlyKhWyMw+H")
	assert.Nil(t, err)

	config = Config{}

	err = LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")
	assert.Nil(t, err)

	devConfig = config["development"].(map[string]interface{})
	assert.Equal(t, "hello", devConfig["secret"])

	err = os.Setenv("GO_ENV", "development")
	assert.Nil(t, err)
}

func TestLoad(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")
	assert.Nil(t, err)

	devConfig := config["development"].(map[string]interface{})
	assert.Equal(t, "hello", devConfig["secret"])

	err = Load("../config/env.yml")
	assert.Nil(t, err)

	devConfig = config["development"].(map[string]interface{})
	assert.Equal(t, "value", devConfig["key"])
	assert.Equal(t, "hello", devConfig["secret"])
}
