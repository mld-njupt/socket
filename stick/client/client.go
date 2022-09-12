package client

import (
	"fmt"
	"net"

	"github.com/mld-njupt/socket/stick/proto"
)

func Client() {
	conn, err := net.Dial("tcp", "127:0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hello, Hello. How are you?"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
