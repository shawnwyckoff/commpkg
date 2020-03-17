package gproxy

import (
	"github.com/pkg/errors"
	// "github.com/GameXG/ProxyClient" // http client using socks5 proxy supported not well
	"github.com/shawnwyckoff/gpkg/dsa/gspeed"
	"github.com/shawnwyckoff/gpkg/net/addr"
	"github.com/shawnwyckoff/gpkg/net/ghttp"
	"time"
)

const (
	proxyDetectURL = "http://www.baidu.com"
)

type ProxyQuality struct {
	Type      string
	Available bool
	Speed     *gspeed.Speed
	Latency   time.Duration
}

func CheckProxy(hostAddr string, t string) (*ProxyQuality, error) {
	if t == "unknown" {
		return nil, errors.New("Unknown proxy type")
	}
	_, _, err := addr.ParseHostAddrOnline(hostAddr)
	if err != nil {
		return nil, err
	}

	var pq ProxyQuality
	pq.Available = false
	var counter *gspeed.SpeedCounter = gspeed.NewCounter(time.Minute)

	if t == "http" || t == "https" || t == "socks5" {
		counter.BeginCount()
		resp, err := ghttp.Get(proxyDetectURL, t+"://"+hostAddr, time.Second*5, true)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != 200 {
			return &pq, nil
		}
		s, err := ghttp.ReadBodyString(resp)
		if err != nil {
			return nil, err
		}
		if len(s) == 0 {
			return nil, errors.New("Empty content")
		}
		pq.Available = true
		counter.Add(uint64(len(resp.Header) + len(s)))
		pq.Speed, err = counter.Get()
		if err != nil {
			return nil, err
		}
		return &pq, nil
	} else {
		return nil, errors.New(t + " type unsupported for now")
	}
}

/*
func TryProxy(hostAddr string) (available bool, t ProxyType, err error) {
	_, _, err = xhostaddr.ParseAddrString(hostAddr)
	if err != nil {
		return false, PROXY_TYPE_UNKNOWN, err
	}

	available, err = CheckProxy(hostAddr, PROXY_TYPE_HTTP)
	if err == nil && available {
		return true, PROXY_TYPE_HTTP, nil
	}
	available, err = CheckProxy(hostAddr, PROXY_TYPE_HTTPS)
	if err == nil && available {
		return true, PROXY_TYPE_HTTPS, nil
	}
	available, err = CheckProxy(hostAddr, PROXY_TYPE_SOCKS5)
	if err == nil && available {
		return true, PROXY_TYPE_SOCKS5, nil
	}
	return false, PROXY_TYPE_UNKNOWN, nil
}*/
