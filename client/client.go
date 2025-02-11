package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server := "localhost:8080"

	conn, err := net.Dial("tcp", server)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Подключено к серверу!")

	message := "Привет, сервер!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при получении ответа:", err)
		os.Exit(1)
	}

	fmt.Printf("Ответ сервера: %s\n", string(buffer[:n]))
}
