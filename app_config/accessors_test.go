package AppConfig

import "testing"

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
