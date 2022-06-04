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

func send_comand(host string, commandRaw string) {
	conn, err := create_dial(host)
	if err != nil {
		log.Println("client: net dial connection create error", err)
		os.Exit(1)
		return
	}
	defer conn.Close()

	err = connection.ConnectionSend("", connection.ConnectionTypeCreate, conn)
	if err != nil {
		log.Println("client: connection create error", err)
		os.Exit(1)
		return
	}

	var idBuffer = make([]byte, 1024)
	_, err = conn.Read(idBuffer)
	if err != err {
		log.Println("client: error read id", err)
		os.Exit(1)
	}
	var id = string(idBuffer)

	command := utils.StringToCommand(commandRaw)
	err = command.Send(conn)
	if err != err {
		log.Println("client: net write command ", err)
		os.Exit(1)
	}

	var waitFinish = new(sync.WaitGroup)
	waitFinish.Add(3)

	// ready readyStart
	var readyStart = make([]byte, 512)
	_, err = conn.Read(readyStart)
	if err != err {
		log.Println("client: error read readyStart ", err)
		os.Exit(1)
	}

	go create_stder(id, host, waitFinish)
	go create_stdin(id, host, waitFinish)
	go create_stdout(id, host, waitFinish)

	// ready readyAfterClose
	_, err = conn.Read(make([]byte, 512))
	if err != err {
		log.Println("client: error read readyAfterClose ", err)
		os.Exit(1)
	}

	// send stdout close
	_, err = conn.Write(make([]byte, 512))
	if err != err {
		log.Println("client: error send close stdout ", err)
		os.Exit(1)
	}

	waitFinish.Wait()
}

func create_stdin(id string, host string, waitFinish *sync.WaitGroup) {
	defer waitFinish.Done()
	conn, err := create_dial(host)
	if err != nil {
		log.Println("client: create stdin err", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeStdin, conn)
	if err != nil {
		log.Println("client: connection create stdin error")
	}
	_, err = conn.Read(make([]byte, 512))
	if err != nil {
		log.Println("client: connection wait send stdin")
	}

	io.Copy(conn, os.Stdin)
}

func create_stdout(id string, host string, waitFinish *sync.WaitGroup) {
	defer waitFinish.Done()
	conn, err := create_dial(host)
	if err != nil {
		log.Println("client: create stdout err", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeStdout, conn)
	if err != nil {
		log.Println("client: connection create stdout error")
	}
	io.Copy(os.Stdout, conn)
}

func create_stder(id string, host string, waitFinish *sync.WaitGroup) {
	defer waitFinish.Done()
	conn, err := create_dial(host)
	if err != nil {
		log.Println("client: create stdout err", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = connection.ConnectionSend(id, connection.ConnectionTypeStder, conn)
	if err != nil {
		log.Println("client: connection create stdout error")
	}
	io.Copy(os.Stderr, conn)
}
