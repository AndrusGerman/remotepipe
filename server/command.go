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
	stder, _ := cmd.StderrPipe()
	stdout, _ := cmd.StdoutPipe()
	stdin, _ := cmd.StdinPipe()

	// send response
	go func() {
		defer stdout.Close()
		io.Copy(ctx.Stdout, stdout)
	}()

	// get pipe
	go func() {
		defer stdin.Close()
		io.Copy(stdin, ctx.Stdin)
	}()

	// send errors
	go func() {
		defer stder.Close()
		io.Copy(ctx.Stder, stder)
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
	log.Println("server: finish comand")
}
