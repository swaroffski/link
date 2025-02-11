package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Создаем TCP-сервер на порту 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании сервера:", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("Сервер запущен на порту 8080")

	// Ожидаем подключения
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		// Чтение данных от клиента
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
			break
		}
		fmt.Println("Получено сообщение:", string(buffer))
	}
	conn.Close()
}
