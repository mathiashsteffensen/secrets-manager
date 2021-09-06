package crypto

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGenRandomBytes(t *testing.T) {
	byteLength := 32

	last, err := GenRandomBytes(byteLength)
	assert.Nil(t, err)

	for i := 0; i < 2_000_000; i++ {
		current, err := GenRandomBytes(byteLength)
		assert.Nil(t, err)
		assert.NotEqualf(t, string(last), string(current), "GenRandomBytes should not return the same thing again in a 2_000_000 iteration loop, failed on %s", strconv.Itoa(i))
	}
}
