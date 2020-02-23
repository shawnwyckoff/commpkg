package probe

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/commpkg/dsa/num"
	"github.com/shawnwyckoff/commpkg/net/addr"
	"net"
	"time"
)

func TcpingOnline(host string, port int, timeout time.Duration) (opened bool, err error) {
	if !addr.IsValidPort(port) {
		return false, errors.Errorf("Invalid port " + num.ToString(port))
	}

	ip, _, err := addr.ParseHostAddrOnline(host)
	if err != nil {
		return false, err
	}

	conn, err := net.DialTimeout("tcp", ip.String()+":"+num.ToString(port), timeout)
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
