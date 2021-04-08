package main

import (
	"fmt"
	"net"
)

// *** UDP编程-服务端 ***
// UDP协议（User Datagram Protocol）中文名称是用户数据报协议，
// 是OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议，
// 不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，
// 但是UDP协议的实时性比较好，通常用于视频直播相关领域。
func main()  {
	// UDP服务端
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
