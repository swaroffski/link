package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
)

// Генерация ключей
func GenerateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return priv, &priv.PublicKey
}

// Вычисление общего секрета
func ComputeSharedSecret(priv *ecdsa.PrivateKey, pub *ecdsa.PublicKey) []byte {
	x, _ := pub.Curve.ScalarMult(pub.X, pub.Y, priv.D.Bytes())
	sharedSecret := sha256.Sum256(x.Bytes())
	return sharedSecret[:]
}

// Пример использования
func main() {
	privA, pubA := GenerateKeys()
	privB, pubB := GenerateKeys()

	sharedSecretA := ComputeSharedSecret(privA, pubB)
	sharedSecretB := ComputeSharedSecret(privB, pubA)

	if string(sharedSecretA) == string(sharedSecretB) {
		fmt.Println("Обмен ключами успешен!")
	} else {
		fmt.Println("Ошибка в обмене ключами.")
	}
}
