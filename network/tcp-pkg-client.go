package main

import (
	"fmt"
	"net"
)

// *** TCP黏包-客户端 ***
func main()  {
	// 客户端分10次发送的数据，在服务端并没有成功的输出10次，而是多条数据“粘”到了一起。
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
	}
}

