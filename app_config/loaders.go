package AppConfig

import (
	"github.com/ghodss/yaml"
	"github.com/ieee0824/go-deepmerge"
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"io/ioutil"
	"path/filepath"
)

var (
	config = map[string]interface{}{}
	ENV = env("GO_ENV", "development")
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

	return nil
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

func loadYaml(yamlBytes []byte, target *map[string]interface{}) (err error) {
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
	newConfig := map[string]interface{}{}

	err = loadYaml(bytes, &newConfig)

	if err != nil {
		return
	}

	configInterface, err := deepmerge.Merge(config, newConfig)

	config = configInterface.(map[string]interface{})

	return
}
