package main

import (
	"flag"

	"github.com/AndrusGerman/remotepipe/client"
	"github.com/AndrusGerman/remotepipe/scan"
	"github.com/AndrusGerman/remotepipe/server"
)

var serverBool = false
var scanBool = false

func init() {
	flag.BoolVar(&serverBool, "server", serverBool, "server mode")
	flag.BoolVar(&scanBool, "scan", scanBool, "scan mode")

	flag.Parse()
}

func main() {

	// scan mode
	if scanBool {
		scan.Execute()
		return
	}

	// server mode
	if serverBool {
		server.Execute()
		return
	}

	// client mode
	if !serverBool {
		client.Execute(flag.Args())
		return
	}
}
