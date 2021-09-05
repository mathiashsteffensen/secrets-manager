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
	"fmt"
	"strings"
)

func Get(keys string) (value interface{}, err error) {
	keysSlice := make([]string, 1, len(keys)+1)
	keysSlice[0] = ENV
	keysSlice = append(keysSlice, strings.Split(keys, ".")...)

	nestedConfig := config

	for i, key := range keysSlice {
		var ok bool
		if i != len(keysSlice) - 1 {
			nestedConfig, ok = nestedConfig[key].(map[string]interface{})
		} else {
			value, ok = nestedConfig[key]
		}

		if !ok {
			err = errors.New(fmt.Sprintf("key not found, key: %s", keys))
			return
		}
	}

	return
}

func GetOrDefault(keys string, defaultValue interface{}) interface{} {
	value, err := Get(keys)
	if err != nil {
		return defaultValue
	}
	return value
}

func MustGet(keys string) (value interface{}) {
	value, err := Get(keys)
	if err != nil {
		panic(err)
	}
	return
}

func Exists(keys string) bool {
	_, err := Get(keys)
	if err != nil {
		return false
	}
	return true
}
