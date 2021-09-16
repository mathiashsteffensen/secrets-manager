package cmd

import (
	"os"
	"testing"
)

func setup() {
	secretsFile = "../config/secrets.yml.enc"
	keyFile = "../config/master.key"
	envFiles = []string{"../config/env.yml"}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func shutdown() {

}
