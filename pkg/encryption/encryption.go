package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plainText, nil)
	return ciphertext, nil
}

func main() {
	// Пример ключа и текста
	key := []byte("thisis32bitlongpassphraseimusing")
	plainText := []byte("Hello, world!")

	ciphertext, err := encrypt(plainText, key)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	fmt.Printf("Зашифрованный текст: %x\n", ciphertext)
}
