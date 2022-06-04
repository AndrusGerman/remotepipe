package utils

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/AndrusGerman/remotepipe/config"
)

func GetDevicesByIP(ip string, onFound func(host string)) {
	var spl = strings.Split(ip, ".")

	baseIPRaw := spl[:len(spl)-1]
	var baseIP = strings.Join(baseIPRaw, ".")

	var wait = new(sync.WaitGroup)

	for i := 1; i < 44; i++ {
		wait.Add(1)
		host := fmt.Sprintf("%s.%d", baseIP, i)
		go func() {
			defer wait.Done()
			onFound(host)
		}()
	}

	wait.Wait()
}

func DialDeviceIsRemotepipeServer(host string) bool {
	d := net.Dialer{Timeout: time.Second * 7}

	conn, err := d.Dial("tcp", host+":"+config.PortTCP)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
