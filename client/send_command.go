package client

import (
	"io"
	"log"
	"net"
	"os"
	"sync"

	"github.com/AndrusGerman/remotepipe/config"
	"github.com/AndrusGerman/remotepipe/pkg/connection"
	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func create_dial(host string) (net.Conn, error) {
	return net.Dial("tcp", host+":"+config.PortTCP)
}

func send_comand(id string, host string, commandRaw string) {
	conn, err := create_dial(host)
	if err != err {
		log.Println("client: net dial connection create error")
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeCreate, conn)
	if err != err {
		log.Println("client: connection create error")
		os.Exit(1)
	}

	var waitFinish = new(sync.WaitGroup)
	waitFinish.Add(3)

	go create_stder(id, host, waitFinish)
	go create_stdin(id, host, waitFinish)
	go create_stdout(id, host, waitFinish)

	command := utils.StringToCommand(commandRaw)
	err = command.Send(conn)
	if err != err {
		log.Println("client: net write command ", err)
		os.Exit(1)
	}

	var finishSignal = make([]byte, 512)
	conn.Read(finishSignal)
	waitFinish.Wait()
}

func create_stdin(id string, host string, waitFinish *sync.WaitGroup) {
	defer waitFinish.Done()
	conn, err := create_dial(host)
	if err != err {
		log.Println("client: create stdin err", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeStdin, conn)
	if err != err {
		log.Println("client: connection create stdin error")
	}
	io.Copy(conn, os.Stdin)
}

func create_stdout(id string, host string, waitFinish *sync.WaitGroup) {
	defer waitFinish.Done()
	conn, err := create_dial(host)
	if err != err {
		log.Println("client: create stdout err", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeStdout, conn)
	if err != err {
		log.Println("client: connection create stdout error")
	}
	io.Copy(os.Stdout, conn)
}

func create_stder(id string, host string, waitFinish *sync.WaitGroup) {
	defer waitFinish.Done()
	conn, err := create_dial(host)
	if err != err {
		log.Println("client: create stdout err", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeStder, conn)
	if err != err {
		log.Println("client: connection create stdout error")
	}
	io.Copy(os.Stderr, conn)
}
