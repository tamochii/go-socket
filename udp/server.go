package main

import (
    "fmt"
    "net"
)

func main() {
    // 1. 解析地址
    addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
    if err != nil {
        fmt.Println("地址解析失败:", err)
        return
    }

    // 2. 监听 UDP (ListenUDP)
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("启动 UDP 服务器失败:", err)
        return
    }
    defer conn.Close()
    fmt.Println("UDP 服务端已启动，正在监听 127.0.0.1:8081 ...")

    // 创建一个缓冲区用于接收数据
    buffer := make([]byte, 1024)

    for {
        // 3. 读取数据 (ReadFromUDP)
        // 返回读取的字节数 n，以及发送者的地址 remoteAddr
        n, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("读取数据失败:", err)
            continue
        }

        message := string(buffer[:n])
        fmt.Printf("收到来自 %s 的消息: %s\n", remoteAddr, message)

        // 4. 发送回执 (WriteToUDP)
        // UDP 是无连接的，所以发送时必须指定目标地址 remoteAddr
        reply := fmt.Sprintf("服务端已收到: %s", message)
        _, err = conn.WriteToUDP([]byte(reply), remoteAddr)
        if err != nil {
            fmt.Println("发送回执失败:", err)
        }
    }
}