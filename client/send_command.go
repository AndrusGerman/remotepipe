package client

import (
	"io"
	"log"
	"net"
	"os"

	"github.com/AndrusGerman/remotepipe/config"
	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func send_comand(host string, commandRaw string) {
	conn, err := net.Dial("tcp", host+":"+config.PortTCP)
	if err != err {
		log.Println("client: net dial connection error")
		os.Exit(1)
	}
	defer conn.Close()

	command := utils.StringToCommand(commandRaw)
	err = command.Send(conn)
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
