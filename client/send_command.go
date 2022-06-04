package client

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/AndrusGerman/remotepipe/config"
)

func send_comand(host string, command string) {
	var buffCommand = make([]byte, config.NetworkCommandSize)
	copy(buffCommand, []byte(command))

	conn, err := net.Dial("tcp", host+":"+config.PortTCP)
	if err != err {
		log.Println("client: net dial connection error")
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write(buffCommand)
	if err != err {
		log.Println("client: net write command ", err)
		os.Exit(1)
	}

	// send pipe
	go func() {
		io.Copy(conn, os.Stdin)
	}()

	// get response
	io.Copy(os.Stdout, conn)
}
