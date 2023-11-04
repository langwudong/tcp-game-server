package main

import (
	"game_server/src/network"
	"github.com/sirupsen/logrus"
	"net"
)

func startTCPServer() {
	listener, err := net.Listen(network.Network, network.Port)
	if err != nil {
		logrus.Error("net.Listen err:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logrus.Error("Error accepting connection:", err)
			continue
		}

		go network.HandleLogin(conn)
	}
}

func main() {
	startTCPServer()
}
