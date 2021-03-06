package server

import (
	"log"
	"net"

	"github.com/AndrusGerman/remotepipe/config"
	"github.com/AndrusGerman/remotepipe/pkg/connection"
)

func start_dial_tcp_server() error {
	server, err := net.Listen("tcp", ":"+config.PortTCP)
	if err != nil {
		return err
	}

	var loop = true

	for loop {
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

	metadata, err := connection.ConnectionRead(client)
	if err != nil {
		log.Println("server: error read dial", err)
		return
	}

	if metadata.Type == connection.ConnectionTypeCreate {
		var id = newID()
		metadata.ID = string(id)
		_, err = client.Write(id)
		if err != nil {
			log.Println("server: error send id", err)
			return
		}

		NewProccess(metadata, client)
		return
	}

	SendToProccess(metadata, client)
}
