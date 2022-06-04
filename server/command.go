package server

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func run_cmd(command []byte, conn net.Conn) {
	var err error
	comand := utils.ByteToCommand(command)

	log.Printf("comand: '%s', flangs'%v'\n", comand.Command, comand.Flags)
	cmd := exec.Command(comand.Command, comand.Flags...)
	cmd.Stderr = os.Stderr
	stdout, _ := cmd.StdoutPipe()
	stdin, _ := cmd.StdinPipe()
	defer stdout.Close()
	defer stdin.Close()

	// send response
	go func() {
		io.Copy(conn, stdout)
	}()

	// get pipe
	go func() {
		io.Copy(stdin, conn)
	}()

	err = cmd.Run()
	if err != nil {
		log.Println("server: comand", err)
	}
}
