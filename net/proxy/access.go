package proxy

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/net/probe"
	"time"
)

func IsVisitable(url string) bool {
	fmt.Println(probe.TcpingOnline("www.youtube.com", 443, time.Millisecond*500))
	return false
}
