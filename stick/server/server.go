package server

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/mld-njupt/socket/stick/proto"
)

// 处理函数
func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read client failed, err", err)
			break
		}
		fmt.Println("client data is:", msg)
	}
}
func Server() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
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
