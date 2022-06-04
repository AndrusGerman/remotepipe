package server

import "log"

func Execute() {
	log.Println("server is start")
	var err = start_dial_tcp_server()
	if err != nil {
		log.Println("server: err start server ", err)
	}
}
