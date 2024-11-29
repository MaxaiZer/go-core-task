package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
)

func printType[T any](value T) {
	fmt.Printf("type %v of value %v\n", reflect.TypeOf(value), value)
}

func appendToStr[T any](str *string, value T) {
	if str == nil {
		*str = ""
	}

	*str += fmt.Sprintf("%v", value)
}

func insert[T any](text *[]T, salt []T, idx int) error {
	if idx < 0 || idx > len(*text) {
		return errors.New("index out of bounds")
	}

	res := make([]T, 0, len(*text)+len(salt))
	res = append(res, (*text)[:idx]...)
	res = append(res, salt...)
	res = append(res, (*text)[idx:]...)
	*text = res
	return nil
}

func aesEncrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	initVector := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, initVector); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, initVector)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], text)
	return ciphertext, nil
}

func aesDecrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	initVector := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, initVector)
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	values := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
	for _, v := range values {
		printType(v)
	}

	all := ""

	for _, v := range values {
		appendToStr(&all, v)
		all += " "
	}

	fmt.Println(all)

	runes := []rune(all)
	fmt.Printf("runes: %v\n", runes)

	bytes := []byte(string(runes))
	err := insert(&bytes, []byte("go-2024"), len(bytes)/2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("before encrypt: " + string(bytes))

	key := []byte("DntStealMySuperDuperSecretKeyPls")
	encrypted, err := aesEncrypt(key, bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("encrypted: " + string(encrypted))

	decrypted, err := aesDecrypt(key, encrypted)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("decrypted: " + string(decrypted))
}
