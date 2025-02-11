package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Запуск сервера на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Слушаем порт 8080...")

	// Ожидаем подключения
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Отправляем сообщение
	conn.Write([]byte("Привет, друг!"))

	// Получаем сообщение
	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	fmt.Printf("Получено сообщение: %s\n", string(buffer[:n]))
}
