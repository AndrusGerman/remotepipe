package scan

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/AndrusGerman/getmyip"
	"github.com/AndrusGerman/remotepipe/config"
)

var found = 0

func Execute() {
	fmt.Println("scan: searching devices...")
	ips := getmyip.GetLocalIP()
	for _, v := range ips {
		getDevicesByIP(v)
	}

	fmt.Println("scan: found ", found)
}

func getDevicesByIP(ip string) {
	var spl = strings.Split(ip, ".")

	baseIPRaw := spl[:len(spl)-1]
	var baseIP = strings.Join(baseIPRaw, ".")

	var wait = new(sync.WaitGroup)

	for i := 1; i < 44; i++ {
		wait.Add(1)
		host := fmt.Sprintf("%s.%d", baseIP, i)
		go dialDeviceScan(host, wait)
	}

	wait.Wait()
}

func dialDeviceScan(host string, wait *sync.WaitGroup) {
	d := net.Dialer{Timeout: time.Second * 7}

	defer wait.Done()
	conn, err := d.Dial("tcp", host+":"+config.PortTCP)
	if err != nil {
		return
	}
	defer conn.Close()
	fmt.Println("scan: " + host)
	found++
}
