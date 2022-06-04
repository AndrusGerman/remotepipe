package server

import (
	"log"
	"net"

	"github.com/AndrusGerman/remotepipe/config"
	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func start_dial_tcp_server() error {
	server, err := net.Listen("tcp", ":"+config.PortTCP)
	if err != nil {
		return err
	}

	var loop = true

	for loop {
		log.Println("server: waiting client")
		client, err := server.Accept()
		if err != nil {
			log.Println("server: client connection fail", err)
			continue
		}
		go client_connections(client)
	}
	return err
}

func client_connections(client net.Conn) {
	defer client.Close()

	var command = new(utils.Command)
	err := command.Get(client)
	if err != nil {
		log.Println("server: error read dial", err)
		return
	}
	run_cmd(command, client)
}
