package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Подключаемся к серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Чтение и отправка сообщения
	fmt.Println("Введите сообщение для отправки на сервер:")
	var message string
	fmt.Scanln(&message)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Ошибка при отправке:", err)
	}
}
