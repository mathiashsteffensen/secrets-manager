package FileHelpers

import (
	"embed"
	"encoding/base64"
	"github.com/mathiashsteffensen/secrets-manager/crypto"
	"os"
	"path/filepath"
)

func Decrypt(secrets []byte, key []byte) (decrypted []byte, err error) {
	decoded, err := base64.StdEncoding.DecodeString(string(secrets))
	if err != nil {
		return
	}

	decrypted, err = crypto.Decrypt(decoded, key)

	return
}

func LoadFile(relativePath string) (contents []byte, err error) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		return
	}

	contents, err = os.ReadFile(absolutePath)
	return
}

func LoadFileFS(fs *embed.FS, path string) (contents []byte, err error) {
	return fs.ReadFile(path)
}
