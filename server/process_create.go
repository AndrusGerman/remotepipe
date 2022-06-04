package server

import (
	"log"
	"net"
	"sync"

	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func (ctx *Proccess) Start() {
	log.Println("server: starting proccess ", ctx.ID)
	defer ctx.Close()
	ctx.InitChannels()
	waitingCanales := ctx.GetCanales()
	var command = new(utils.Command)
	err := command.Get(ctx.Create)
	if err != nil {
		log.Println("server: error read dial", err)
		return
	}

	_, err = ctx.Create.Write(make([]byte, 512)) // readyStart signal send
	if err != nil {
		log.Println("server: error read readyStart", err)
		return
	}

	log.Println("server: waiting channels")

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

	if ctx.Stder != nil {
		log.Println("close 1: send", ctx.Stder.Close())
	}

	if ctx.Stdin != nil {
		log.Println("close 2: send", ctx.Stdin.Close())
	}

	if ctx.Stdout != nil {
		_, err := ctx.Create.Read(make([]byte, 512)) // close stdout signal
		log.Println("close 3: read", err)
		ctx.Stdout.Close()
	}

	if ctx.Create != nil {
		log.Println("close 4: send", ctx.Create.Close())
	}

	delete(proccess, ctx.ID)
	log.Println("server: close ", ctx.ID)
}
