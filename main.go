package main

import (
	"flag"

	"github.com/AndrusGerman/remotepipe/client"
	"github.com/AndrusGerman/remotepipe/server"
)

var serverBool = false

func init() {
	flag.BoolVar(&serverBool, "server", serverBool, "server mode")
	flag.Parse()
}

func main() {

	// server mode
	if serverBool {
		server.Execute()
	}

	// client mode
	if !serverBool {
		client.Execute(flag.Args())
	}
}
