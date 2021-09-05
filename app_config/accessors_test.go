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
	"testing"
)

func TestGet(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")

	if err != nil {
		t.Errorf("Unexpected error when loading encrypted secrets %#v", err)
	}

	got, err := Get("secret")

	if err != nil {
		t.Errorf("Unexpected error when running Get %#v", err)
	}

	if got.(string) != "hello" {
		t.Errorf(`Expected Get("secret") to equal "hello" but got %#v`, got.(string))
	}

	got, err = Get("super.deeply.nested")

	if err != nil {
		t.Errorf("Unexpected error when running Get %#v", err)
	}

	if got.(string) != "value" {
		t.Errorf(`Expected Get("super.deeply.nested") to equal "value" but got %#v`, got.(string))
	}
}

func TestExists(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")

	if err != nil {
		t.Errorf("Unexpected error when loading encrypted secrets %#v", err)
	}

	got := Exists("secret")

	if got != true {
		t.Errorf("Expected key 'secret' to exists but it didn't")
	}

	got = Exists("this.is.not.a.real.key")

	if got != false {
		t.Errorf("Expected key 'secret' to not exists but it did")
	}
}

func TestAllKeys(t *testing.T) {
	err := LoadEncrypted("../config/secrets.yml.enc", "../config/master.key")

	if err != nil {
		t.Errorf("Unexpected error when loading encrypted secrets %#v", err)
	}

	for _, s := range AllKeys() {
		if !Exists(s) {
			t.Errorf("Expected key %s to exist but it didn't", s)
		}
	}
}
