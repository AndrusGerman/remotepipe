package client

import (
	"log"
	"os"

	"github.com/AndrusGerman/getmyip"
	"github.com/AndrusGerman/remotepipe/config"
	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

func parse_host(host string) string {
	if host == config.HostFind {
		return getFirstHost()
	}
	return host
}

func getFirstHost() string {
	var ipBff = make(chan string, 50)
	ips := getmyip.GetLocalIP()
	go func() {
		for _, v := range ips {
			GetDevicesByIP(v, ipBff)
		}
		defer close(ipBff)
	}()

	first := <-ipBff
	if first == "" {
		log.Println("client: ip default not found")
		os.Exit(1)
	}
	return first
}

func GetDevicesByIP(ip string, ipBff chan string) {
	utils.GetDevicesByIP(ip, func(host string) {
		var b = utils.DialDeviceIsRemotepipeServer(host)
		if b {
			ipBff <- host
		}
	})
	// send none, default
	ipBff <- ""
}
