package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func Encrypt(data []byte, key string) []byte {
	aesBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, data, nil)

	return cipheredText
}

func Decrypt(data []byte, key string) []byte {
	aesBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := data[:nonceSize], data[nonceSize:]

	plainText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		fmt.Println(err)
	}

	return plainText
}
