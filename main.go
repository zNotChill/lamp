package main

import (
	"fmt"
	"net"

	"znci.dev/lamp-v2/connection"
	"znci.dev/lamp-v2/utils"
)

func main() {
	config := utils.LoadConfig()

	listener, err := net.Listen("tcp", fmt.Sprint(":", config.Port))
	if err != nil {
		utils.Error("Failed to start server")
	}

	utils.Info("Server started on port " + fmt.Sprint(config.Port))

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.Error("Failed to accept connection")
		}

		go connection.HandleConnection(conn, connection.ProtocolClient{})
	}
}
