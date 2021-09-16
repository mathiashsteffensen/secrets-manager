package FileHelpers

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadEncryptedSecretsFile(t *testing.T) {
	key, err := ioutil.ReadFile("../config/master.key")
	assert.Nil(t, err)

	contents, err := ReadEncryptedSecretsFile("../config/secrets.yml.enc", key)
	assert.Nil(t, err)

	expectedContents := []byte("production:\n  key: other-value\n  super:\n    deeply:\n      nested: value\ndevelopment:\n  secret: hello\n  super:\n    deeply:\n      nested: value")

	assert.Equal(t, expectedContents, contents)
}
