package server

import (
	"bufio"
	"fmt"
	"net"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read client failed, err", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("client data is:", recvStr)
		conn.Write([]byte(recvStr))
	}
}
func Server() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("server listen, err", err)
		return
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accept failed, err", err)
			continue
		}
		go process(conn)
	}
}
