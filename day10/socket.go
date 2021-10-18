package main

/**
 * socket 网络编程
 */

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

var exit = true
var protocol = "tcp"
var address = "127.0.0.1:8080"

func handler(conn net.Conn) string {
	reader := bufio.NewReader(conn)
	var buf [256]byte
	n, err := reader.Read(buf[:])
	if err != nil {
		errors.New("数据读取失败")
	}
	return string(buf[:n])
}

func serverHandler(conn net.Conn) {
	for {
		msg := handler(conn)
		fmt.Println("接收客户端发送到数据: ", msg)
		_, err := conn.Write([]byte("[from server]: " + time.Now().String()))
		if err != nil {
			return
		}
	}
}

func clientHandler(conn net.Conn) {
	for {
		msg := handler(conn)
		fmt.Println("接收服务端响应数据: ", msg)
	}
}

// 服务端
func startServer() {
	// 1. 监听8080端口, 创建一个socket
	serverSocket, err := net.Listen(protocol, address)
	if err != nil {
		fmt.Println("Listen ", address, " 失败, err: ", err)
		return
	}
	for {
		// 死循环获取连接
		conn, err := serverSocket.Accept()
		if err != nil {
			fmt.Println("Accept 失败, err: ", err)
			continue
		}
		go serverHandler(conn)
	}
}

// 客户端
func startClient() {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		fmt.Println("Dial ", address, " 失败, err: ", err)
		return
	}
	for {
		var msg string
		fmt.Scan(&msg)
		if msg == "exit" {
			exit = false
			fmt.Println("退出啦~~~")
			return
		}
		_, err = conn.Write([]byte("[from client]: " + msg))
		if err != nil {
			return
		}
		go clientHandler(conn)
	}
}

func main() {
	go startServer()
	go startClient()

	for exit {
		time.Sleep(time.Second)
	}
}
