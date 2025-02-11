package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"link/pkg/dict"     // Пакет для словаря
	"link/pkg/user"     // Пакет для пользователей
	"link/internal/db"  // Пакет для базы данных
)

func main() {
	// Генерация уникального логина
	login := user.GenerateLogin()
	fmt.Println("Ваш уникальный логин:", login)

	// Генерация ключевых слов
	words := dict.LoadWords("pkg/dict/words.txt") // Загружаем слова из файла
	key := user.GenerateKey(words)                 // Генерируем ключ по словам

	// Сопоставление логина и открытого ключа
	db.SaveUser(login, key) // Сохраняем данные пользователя в базу
	fmt.Println("Ключ для доступа:", key)

	// Печать всех пользователей
	db.PrintUsers()