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
	"fmt"
	"strings"
)

// Get a specified key from your configuration
func Get[TReturns any](keys string) (value TReturns, err error) {
	keysSlice := make([]string, 1, len(keys)+1)
	keysSlice[0] = ENV
	keysSlice = append(keysSlice, strings.Split(keys, ".")...)

	nestedConfig := config

	var untypedValue any

	for i, key := range keysSlice {
		isLastKey := i == len(keysSlice)-1

		var ok bool
		if !isLastKey {
			nestedConfig, ok = nestedConfig[key].(Config)
		} else {
			untypedValue, ok = nestedConfig[key]
		}

		if !ok {
			err = fmt.Errorf("AppConfig: key not found in loaded configuration, key: %s", keys)
			return
		}

		if isLastKey {
			value, ok = untypedValue.(TReturns)

			if !ok {
				err = fmt.Errorf("AppConfig: type assertion failed for key: %s", keys)
				return
			}
		}
	}

	return
}

// GetOrDefault calls Get and if any errors occur it returns the provided default value instead
func GetOrDefault[TReturns any](keys string, defaultValue TReturns) TReturns {
	value, err := Get[TReturns](keys)
	if err != nil {
		return defaultValue
	}
	return value
}

// MustGet calls Get and panics if any errors occur
func MustGet[TReturns any](keys string) (value TReturns) {
	value, err := Get[TReturns](keys)
	if err != nil {
		panic(err)
	}
	return
}

// Exists returns true if a specified key exists and false otherwise
func Exists(keys string) bool {
	_, err := Get[any](keys)
	return err == nil
}

// GetConfig returns the entire Config map for the current environment
func GetConfig() Config {
	return config[ENV].(Config)
}

// AllKeys returns a slice of all the keys which exists in the current environments Config map
func AllKeys() []string {
	return keysInConfig(config[ENV].(Config), "")
}

// keysInConfig recursively fetches all keys in a Config map
func keysInConfig(c Config, prefix string) []string {
	keys := make([]string, 0)
	for key, value := range c {
		nestedConfig, ok := value.(Config)

		var newPrefix string
		if prefix == "" {
			newPrefix = key
		} else {
			newPrefix = fmt.Sprintf("%s.%s", prefix, key)
		}

		if !ok {
			keys = append(keys, newPrefix)
		} else {
			keys = append(keys, keysInConfig(nestedConfig, newPrefix)...)
		}
	}

	return keys
}
