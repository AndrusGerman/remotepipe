package server

import (
	"log"

	"github.com/AndrusGerman/getmyip"
)

func Execute() {
	log.Println("server: is start")
	log.Println("server: ip ", getmyip.GetLocalIP())
	var err = start_dial_tcp_server()
	if err != nil {
		log.Println("server: err start server ", err)
	}
}
