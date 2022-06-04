package server

import (
	"io"
	"log"
	"os/exec"

	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func (ctx *Proccess) run_cmd(comand *utils.Command) {
	var err error
	log.Printf("comand: '%s', flags: '%v'\n", comand.Command, comand.Flags)
	cmd := exec.Command(comand.Command, comand.Flags...)
	stder, err := cmd.StderrPipe()
	if err != nil {
		log.Println("server: err create pipeStderr")
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("server: err create pipeStdout")
		return
	}
	defer stdout.Close()
	defer stder.Close()

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println("server: err create pipeStdin")
		return
	}

	// get pipe
	go func() {
		defer stdin.Close()
		io.Copy(stdin, ctx.Stdin)
	}()

	// send errors
	go func() {
		io.Copy(ctx.Stder, stder)
	}()

	// send response
	go func() {
		io.Copy(ctx.Stdout, stdout)
	}()

	err = cmd.Start()
	if err != nil {
		log.Println("server: start command err", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		log.Println("server: Wait command err", err)
		return
	}

	var readyAfterClose = make([]byte, 512)
	_, err = ctx.Create.Write(readyAfterClose)
	if err != nil {
		log.Println("server: error read readyAfterClose", err)
		return
	}
	log.Println("server: finish comand")
}
