package dialers

import "net"

type Dialer interface {
	Dial(remoteAddr string) (net.Conn, error)
}
