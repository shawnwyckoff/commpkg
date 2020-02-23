package listeners

import "net"

type Listener interface {
	Listen(listenAddr string) (net.Listener, error)
}
