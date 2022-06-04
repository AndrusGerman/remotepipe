package scan

import (
	"fmt"

	"github.com/AndrusGerman/getmyip"
	"github.com/AndrusGerman/remotepipe/pkg/utils"
)

var found = 0

func Execute() {
	fmt.Println("scan: searching devices...")
	ips := getmyip.GetLocalIP()
	for _, v := range ips {
		GetDevicesByIP(v)
	}
	fmt.Println("scan: found ", found)
}

func GetDevicesByIP(ip string) {
	utils.GetDevicesByIP(ip, func(host string) {
		var b = utils.DialDeviceIsRemotepipeServer(host)
		if b {
			found++
			fmt.Println("scan: " + host)
		}
	})
}
