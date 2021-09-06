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

import "testing"

func TestLoadYaml(t *testing.T) {
	target := map[string]interface{}{}

	err := loadYaml([]byte("a: b\nb: \n  c: d"), &target)

	if err != nil {
		t.Errorf("Unexpected error when loading YAML %#v", err)
	}

	if target["a"] != "b" {
		t.Errorf(`Expected target["a"] to equal "b" but got %#v`, target["a"])
	}

	nested, ok := target["b"].(map[string]interface{})

	if !ok {
		t.Errorf(`Expected target["b"] to be of type map[string]interface{}`)
	}

	if nested["c"] != "d" {
		t.Errorf(`Expected target["b"] to equal "c" but got %#v`, target["b"])
	}
}

func TestLoadEncrypted(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")

	if err != nil {
		t.Errorf("Unexpected error when loading encrypted secrets %#v", err)
	}

	devConfig := config["development"].(map[string]interface{})

	if devConfig["secret"] != "hello" {
		t.Errorf(`Expected devConfig["secret"] to equal "hello" but got %#v`, devConfig["secret"])
	}
}

func TestLoad(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")

	if err != nil {
		t.Errorf("Unexpected error when loading encrypted secrets %#v", err)
	}

	devConfig := config["development"].(map[string]interface{})

	if devConfig["secret"] != "hello" {
		t.Errorf(`Expected devConfig["secret"] to equal "hello" but got %#v`, devConfig["secret"])
	}

	prodConfig := config["production"].(map[string]interface{})

	if prodConfig["secret"] != "world" {
		t.Errorf(`Expected prodConfig["secret"] to equal "world" but got %#v`, prodConfig["secret"])
	}

	err = Load("../config/env.yml")

	if err != nil {
		t.Errorf("Unexpected error when loading yaml file %#v", err)
	}

	devConfig = config["development"].(map[string]interface{})

	if devConfig["key"] != "value" {
		t.Errorf(`Expected devConfig["key"] to equal "value" but got %#v`, devConfig["key"])
	}

	if devConfig["secret"] != "hello" {
		t.Errorf(`Expected devConfig["secret"] to equal "hello" but got #%v`, devConfig["secret"])
	}

	if prodConfig["secret"] != "world" {
		t.Errorf(`Expected prodConfig["secret"] to equal "world" but got %#v`, prodConfig["secret"])
	}
}

func TestLoadFile(t *testing.T) {
	contents, err := loadFile("../config/env.yml")

	if err != nil {
		t.Errorf("Unexpected error when loading file %#v", err)
	}

	expectedContents := `production:
  key: other-value
  secret: hello
development:
  key: value`

	if string(contents) != expectedContents {
		t.Errorf("Expected file contents to be %s but got %s", expectedContents, string(contents))
	}
}
