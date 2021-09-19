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
	"errors"
	"github.com/ghodss/yaml"
	"github.com/ieee0824/go-deepmerge"
	FileHelpers "github.com/mathiashsteffensen/secrets-manager/file_helpers"
)

// Config is an alias for map[string]interface{}
type Config = map[string]interface{}

var (
	config           = Config{}
	ENV              = env("GO_ENV", "development")
	ErrNoGoMasterKey = errors.New("GO_ENV is set to production but no GO_MASTER_KEY is set, not loading encrypted secrets file")
)

func LoadEncrypted(secretsLocation string, keyLocation string) (err error) {
	var key []byte

	if env("GO_ENV", "development") == "production" {
		keyString := env("GO_MASTER_KEY", "")
		if keyString == "" {
			return ErrNoGoMasterKey
		}
		key = []byte(keyString)
	} else {
		key, err = FileHelpers.LoadFile(keyLocation)
		if err != nil {
			return
		}
	}

	decrypted, err := FileHelpers.ReadEncryptedSecretsFile(secretsLocation, key)
	if err != nil {
		return
	}

	err = mergeConfig(decrypted)

	return
}

func Load(files ...string) error {
	for _, file := range files {
		yamlBytes, err := FileHelpers.LoadFile(file)
		if err != nil {
			return err
		}

		err = mergeConfig(yamlBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

func mergeConfig(bytes []byte) (err error) {
	newConfig := Config{}

	err = yaml.Unmarshal(bytes, &newConfig)

	if err != nil {
		return
	}

	configInterface, err := deepmerge.Merge(config, newConfig)

	config = configInterface.(map[string]interface{})

	return
}
