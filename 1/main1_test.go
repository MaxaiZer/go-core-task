package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt(t *testing.T) {

	bytes := []byte("blablablablablablabla")
	bytesLen := len(bytes)
	salt := []byte("go-2024")
	err := insert(&bytes, salt, len(bytes)/2)
	assert.NoError(t, err)
	assert.Equal(t, bytesLen+len(salt), len(bytes))

	key := []byte("DntStealMySuperDuperSecretKeyPls")
	encrypted, err := aesEncrypt(key, bytes)
	assert.NoError(t, err)

	decrypted, err := aesDecrypt(key, encrypted)
	assert.NoError(t, err)
	assert.Equal(t, decrypted, bytes)
}
