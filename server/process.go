package server

import (
	"log"
	"net"

	"github.com/AndrusGerman/remotepipe/pkg/connection"
)

type Proccess struct {
	ID         string
	ChanStdout chan net.Conn
	ChanStder  chan net.Conn
	ChanStdin  chan net.Conn

	Stdout net.Conn
	Stder  net.Conn
	Stdin  net.Conn
	Create net.Conn
}

var proccess = make(map[string]*Proccess)

func NewProccess(data *connection.ConnectionMetadata, conn net.Conn) {
	_, exist := proccess[data.ID]
	if exist {
		log.Println("server-CreateProccess: this proccess exist")
		conn.Close()
		return
	}
	var newprocess = new(Proccess)
	proccess[data.ID] = newprocess
	newprocess.ID = data.ID
	newprocess.Create = conn
	newprocess.Start()
}

func SendToProccess(data *connection.ConnectionMetadata, conn net.Conn) {
	if data.ID == "" {
		log.Printf("server-SendToProccess: id is null '%s'\n", data.Type)
		return
	}

	_, exist := proccess[data.ID]
	if !exist {
		log.Printf("server-SendToProccess: this proccess do not exist '%s' '%d'\n'%s'\n", data.ID, len(data.ID), data.Type)
		conn.Close()
		return
	}

	if data.Type == connection.ConnectionTypeStdout {
		proccess[data.ID].ChanStdout <- conn
		return
	}
	if data.Type == connection.ConnectionTypeStdin {
		proccess[data.ID].ChanStdin <- conn
		return
	}
	if data.Type == connection.ConnectionTypeStder {
		proccess[data.ID].ChanStder <- conn
		return
	}
}
