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

func MustGet(keys string) (value interface{}) {
	value, err := Get(keys)
	if err != nil {
		panic(err)
	}
	return
}
