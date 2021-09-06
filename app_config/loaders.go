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
	"github.com/ghodss/yaml"
	"github.com/ieee0824/go-deepmerge"
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"io/ioutil"
	"path/filepath"
)

type Config = map[string]interface{}

var (
	config = Config{}
	ENV    = env("GO_ENV", "development")
)

func LoadEncrypted(secretsLocation string, keyLocation string) (err error) {
	secrets, err := loadFile(secretsLocation)
	if err != nil {
		return
	}

	key, err := loadFile(keyLocation)
	if err != nil {
		return
	}

	decrypted, err := crypto.DecryptSecrets(secrets, key)
	if err != nil {
		return
	}

	err = mergeConfig(decrypted)

	return
}

func Load(files ...string) (err error) {
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
	return
}

func loadYaml(yamlBytes []byte, target *Config) (err error) {
	return yaml.Unmarshal(yamlBytes, target)
}

func loadFile(relativePath string) (contents []byte, err error) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return
	}

	contents, err = ioutil.ReadFile(absolutePath)
	return
}

func mergeConfig(bytes []byte) (err error) {
	newConfig := Config{}

	err = loadYaml(bytes, &newConfig)

	if err != nil {
		return
	}

	configInterface, err := deepmerge.Merge(config, newConfig)

	config = configInterface.(map[string]interface{})

	return
}
