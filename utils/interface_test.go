package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceToKey(t *testing.T) {
	assert.Equal(t, "0x2a", InterfaceToKey(uint8(42)))
	assert.Equal(t, "0x539", InterfaceToKey(uint16(1337)))
	assert.Equal(t, "0x539", InterfaceToKey(uint32(1337)))
	assert.Equal(t, "0x539", InterfaceToKey(uint64(1337)))

	assert.Equal(t, "42", InterfaceToKey(int8(42)))
	assert.Equal(t, "1337", InterfaceToKey(int16(1337)))
	assert.Equal(t, "1337", InterfaceToKey(int32(1337)))
	assert.Equal(t, "1337", InterfaceToKey(int64(1337)))

	assert.Equal(t, "\"foo\"", InterfaceToKey(string("foo")))
	assert.Equal(t, "[]byte{0x66, 0x6f, 0x6f}",
		InterfaceToKey([]byte{'f', 'o', 'o'}))
}
