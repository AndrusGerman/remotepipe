package server

import (
	"log"
	"net"
	"sync"

	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func (ctx *Proccess) Start() {
	defer ctx.Close()
	ctx.InitChannels()
	waitingCanales := ctx.GetCanales()
	var command = new(utils.Command)
	err := command.Get(ctx.Create)
	if err != nil {
		log.Println("server: error read dial", err)
		return
	}

	waitingCanales.Wait()
	log.Println("server: is ready ", ctx.ID)

	ctx.run_cmd(command)
}

func (ctx *Proccess) GetCanales() *sync.WaitGroup {
	var waiting = new(sync.WaitGroup)
	waiting.Add(3)
	go func() {
		ctx.Stder = <-ctx.ChanStder
		waiting.Done()
	}()

	go func() {
		ctx.Stdin = <-ctx.ChanStdin
		waiting.Done()

	}()

	go func() {
		ctx.Stdout = <-ctx.ChanStdout
		waiting.Done()
	}()
	return waiting
}
func (ctx *Proccess) InitChannels() {
	ctx.ChanStder = make(chan net.Conn)

	ctx.ChanStdin = make(chan net.Conn)

	ctx.ChanStdout = make(chan net.Conn)
}

func (ctx *Proccess) Close() {
	close(ctx.ChanStder)
	close(ctx.ChanStdin)
	close(ctx.ChanStdout)

	if ctx.Create != nil {
		ctx.Create.Write([]byte("ok"))
		ctx.Create.Close()
	}
	if ctx.Stder != nil {
		ctx.Stder.Close()
	}
	if ctx.Stdout != nil {
		ctx.Stdout.Close()
	}
	if ctx.Stdin != nil {
		ctx.Stdin.Close()
	}

	delete(proccess, ctx.ID)
	log.Println("server: close ", ctx.ID)
}
