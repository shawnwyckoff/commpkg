package gprobe

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gopkg/container/gnum"
	"github.com/shawnwyckoff/gopkg/net/gaddr"
	"net"
	"time"
)

func TcpingOnline(host string, port int, timeout time.Duration) (opened bool, err error) {
	if !gaddr.IsValidPort(port) {
		return false, errors.Errorf("Invalid port " + gnum.ToString(port))
	}

	ip, _, err := gaddr.ParseHostAddrOnline(host)
	if err != nil {
		return false, err
	}

	conn, err := net.DialTimeout("tcp", ip.String()+":"+gnum.ToString(port), timeout)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		return false, nil // Maybe Closed
	}
	return true, nil // Opened
}
