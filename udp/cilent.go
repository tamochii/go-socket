package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // 1. 解析服务器地址
    serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
    if err != nil {
        fmt.Println("地址解析失败:", err)
        return
    }

    // 2. 建立连接 (DialUDP)
    // 注意：UDP 是无连接的，这里的 Dial 只是在本地初始化一个 socket 并绑定目标地址
    conn, err := net.DialUDP("udp", nil, serverAddr)
    if err != nil {
        fmt.Println("连接失败:", err)
        return
    }
    defer conn.Close()

    inputReader := bufio.NewReader(os.Stdin)
    buffer := make([]byte, 1024)

    for {
        fmt.Print("请输入 UDP 消息: ")
        input, _ := inputReader.ReadString('\n')

        // 3. 发送数据
        _, err := conn.Write([]byte(input))
        if err != nil {
            fmt.Println("发送失败:", err)
            break
        }

        // 4. 接收回执
        n, _, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("接收回执失败:", err)
            break
        }
        fmt.Println("服务器回复:", string(buffer[:n]))
    }
}