package httpz

import (
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// fasthttp

// EnableHTTPProfiling helps easy wake up built in http profiler
func ListenAndServeNoWait(addr string) error {
	if addr == "" {
		return errors.New("nil address")
	}

	var e = make(chan error)
	go func() {
		e <- http.ListenAndServe(addr, nil)
	}()
	select {
	case err := <-e:
		return err
	case <-time.After(time.Millisecond * 5):
		return nil
	}
}
