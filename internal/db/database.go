package db

import (
	"fmt"
)

var users = make(map[string]string) // Храним логины и ключи

// Функция для сохранения пользователя
func SaveUser(login string, key string) {
	users[login] = key
}

// Функция для вывода всех пользователей
func PrintUsers() {
	fmt.Println("Все пользователи:")
	for login, key := range users {
		fmt.Printf("Логин: %s, Ключ: %s\n", login, key)
	}
}
