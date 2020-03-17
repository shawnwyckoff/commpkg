package gprobe

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/dsa/gnum"
	"github.com/shawnwyckoff/gpkg/net/addr"
	"net"
	"time"
)

func TcpingOnline(host string, port int, timeout time.Duration) (opened bool, err error) {
	if !addr.IsValidPort(port) {
		return false, errors.Errorf("Invalid port " + gnum.ToString(port))
	}

	ip, _, err := addr.ParseHostAddrOnline(host)
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
