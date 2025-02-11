package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Шифрование
func encrypt(plainText string, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Создаем новый блок для IV (Initialization Vector)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return cipherText, nil
}

// Дешифрование
func decrypt(cipherText []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}

func main() {
	// Пример использования
	key := []byte("mysecretkey12345") // Ключ шифрования (16 байт для AES-128)
	plainText := "Это секретное сообщение!"

	// Шифруем
	cipherText, err := encrypt(plainText, key)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}
	fmt.Printf("Зашифрованное сообщение: %x\n", cipherText)

	// Дешифруем
	decryptedText, err := decrypt(cipherText, key)
	if err != nil {
		fmt.Println("Ошибка при дешифровании:", err)
		return
	}
	fmt.Printf("Дешифрованное сообщение: %s\n", decryptedText)
}
