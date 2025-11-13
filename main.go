package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	var mode string
	if len(os.Args) < 2 {
		fmt.Print("请选择运行模式 (server/client): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		mode = strings.TrimSpace(input)
	} else {
		mode = os.Args[1]
	}

	if strings.ToLower(mode) == "server" {
		runServer()
	} else if strings.ToLower(mode) == "client" {
		runClient()
	} else {
		fmt.Println("未知的模式:", mode)
		fmt.Println("用法: go run main.go <server|client>")
		os.Exit(1)
	}
}

func runServer() {
	// 监听连接
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("监听错误:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("正在监听 " + CONN_HOST + ":" + CONN_PORT)

	for {
		// 接受一个连接
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("接受连接错误: ", err.Error())
			continue
		}
		// 在新的 goroutine 中处理连接
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("收到来自 %s 的连接\n", conn.RemoteAddr().String())

	for {
		// 从连接中读取数据
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("读取错误:", err.Error())
			}
			break
		}

		// 去除换行符并打印消息
		msg := strings.TrimSpace(string(message))
		fmt.Printf("收到消息: %s\n", msg)

		// 如果客户端发送 "STOP"，则关闭连接
		if msg == "STOP" {
			break
		}

		// 发送响应回客户端
		conn.Write([]byte("消息已收到: " + msg + "\n"))
	}
	fmt.Printf("与 %s 的连接已关闭\n", conn.RemoteAddr().String())
}

func runClient() {
	// 连接到服务器
	conn, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("连接错误:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("已连接到服务器。输入消息并按 Enter 发送。输入 'STOP' 退出。")

	reader := bufio.NewReader(os.Stdin)
	for {
		// 从标准输入读取文本
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')

		// 发送文本到服务器
		fmt.Fprintf(conn, text)

		// 从服务器读取响应
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("从服务器读取错误:", err.Error())
			break
		}
		fmt.Print("来自服务器的消息: " + message)

		// 如果发送了 "STOP"，则退出循环
		if strings.TrimSpace(text) == "STOP" {
			break
		}
	}
	fmt.Println("客户端关闭。")
}
