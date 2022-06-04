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

	log.Printf("comand: '%s', flags: '%v', flagsNumber: '%d' \n", comand.Command, comand.Flags, len(comand.Flags))
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

	err = cmd.Start()
	if err != nil {
		log.Println("server: start command err", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		log.Println("server: Wait command err", err)
	}
}
