package main

import (
	"github.com/mld-njupt/socket/tcp/client"
	"github.com/mld-njupt/socket/tcp/server"
)

func main() {
	go server.Server()
	client.Client()
}
