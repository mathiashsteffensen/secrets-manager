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
	"embed"
	"errors"
	"github.com/ghodss/yaml"
	"github.com/ieee0824/go-deepmerge"
	FileHelpers "github.com/mathiashsteffensen/secrets-manager/file_helpers"
)

type Config = map[string]any

var (
	config           = Config{}
	ENV              = env("GO_ENV", "development")
	ErrNoGoMasterKey = errors.New("GO_ENV is set to production but no GO_MASTER_KEY is set, not loading encrypted secrets file")

	fs *embed.FS
)

func SetFS(newFs *embed.FS) {
	fs = newFs
}

func loadFile(path string) ([]byte, error) {
	if fs == nil {
		return FileHelpers.LoadFile(path)
	} else {
		return FileHelpers.LoadFileFS(fs, path)
	}
}

func LoadEncrypted(secretsLocation string, keyLocation string) (err error) {
	key, err := getKey(keyLocation)
	if err != nil {
		return
	}

	secrets, err := loadFile(secretsLocation)
	if err != nil {
		return
	}

	decrypted, err := FileHelpers.Decrypt(secrets, key)
	if err != nil {
		return
	}

	err = mergeConfig(decrypted)

	return
}

func Load(files ...string) error {
	for _, file := range files {
		yamlBytes, err := loadFile(file)
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

func getKey(keyLocation string) ([]byte, error) {
	var key []byte

	keyString := env("GO_MASTER_KEY", "")

	if keyString == "" {
		if env("GO_ENV", "development") == "production" {
			return key, ErrNoGoMasterKey
		}
		return loadFile(keyLocation)
	} else {
		return []byte(keyString), nil
	}
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
