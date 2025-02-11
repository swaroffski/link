package p2p

import (
	"fmt"
	"net"
)

// Функция для старта сервера (принимает сообщения)
func StartServer(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}
	defer ln.Close()

	fmt.Println("Ожидаем подключения...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения: ", err)
			continue
		}

		// Обработаем подключение
		go handleConnection(conn)
	}
}

// Обработка входящих соединений
func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Новое подключение...")

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении данных: ", err)
		return
	}

	fmt.Printf("Получено сообщение: %s\n", string(buffer[:n]))
	conn.Write([]byte("Сообщение получено!"))
}
