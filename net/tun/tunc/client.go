package tunc

import (
	"fmt"
	"github.com/shawnwyckoff/commpkg/apputil/logger"
	"github.com/shawnwyckoff/commpkg/net/dialers"
	"github.com/shawnwyckoff/commpkg/net/mux"
	"github.com/shawnwyckoff/commpkg/net/smux"
	"github.com/shawnwyckoff/commpkg/sys/ios"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type ClientOption struct {
	LocalTCPListen      string
	RemoteTarget        string
	ScavengeTTL         int
	AutoExpireSeconds   time.Duration
	MuxBufferSize       int
	MuxKeepAliveSeconds int
}

// A pool for stream copying
var xmitBuf sync.Pool

func handleClient(muxer mux.Mux, p1 net.Conn, lg logger.Logger) {
	defer p1.Close()
	p2, err := muxer.Open()
	if err != nil {
		lg.Erro(err)
		return
	}

	defer p2.Close()

	if s2, ok := p2.(mux.Stream); ok {
		lg.Infof("stream opened in: %s out: %s", p1.RemoteAddr(), fmt.Sprint(s2.RemoteAddr(), "(", s2.ID(), ")"))
		defer lg.Infof("stream closed in: %s out: %s", p1.RemoteAddr(), fmt.Sprint(s2.RemoteAddr(), "(", s2.ID(), ")"))
	}

	// start tunnel & wait for tunnel termination
	streamCopy := func(dst io.Writer, src io.ReadCloser) chan struct{} {
		die := make(chan struct{})
		go func() {
			buf := xmitBuf.Get().([]byte)
			if _, err := ios.CopyBuffer(dst, src, buf); err != nil {
				if s2, ok := p2.(mux.Stream); ok {
					// verbose error handling
					cause := err
					if e, ok := err.(interface{ Cause() error }); ok {
						cause = e.Cause()
					}

					if smux.ConvertInternalError(cause) == mux.ErrInvalidProtocol {
						lg.Errof("mux error %s in: %s out: %s", err, p1.RemoteAddr(), fmt.Sprint(s2.RemoteAddr(), "(", s2.ID(), ")"))
					}
				}
			}
			xmitBuf.Put(buf)
			close(die)
		}()
		return die
	}

	select {
	case <-streamCopy(p1, p2):
	case <-streamCopy(p2, p1):
	}
}

func ServeWait(dialer dialers.Dialer, lg logger.Logger, config ClientOption) error {
	rand.Seed(int64(time.Now().Nanosecond()))
	xmitBuf.New = func() interface{} {
		return make([]byte, 4096)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.LocalTCPListen)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	createConn := func() (mux.Mux, error) {
		myconn, err := dialer.Dial(config.RemoteTarget)
		if err != nil {
			return nil, errors.Wrap(err, "dial()")
		}

		// stream multiplex
		session, err := smux.NewSmuxClient(myconn, config.MuxBufferSize, config.MuxKeepAliveSeconds)
		if err != nil {
			return nil, errors.Wrap(err, "createConn()")
		}
		return session, nil

	}

	// wait until a connection is ready
	waitConn := func() mux.Mux {
		for {
			if session, err := createConn(); err == nil {
				return session
			} else {
				lg.Erro(err, "re-connecting:")
				time.Sleep(time.Second)
			}
		}
	}

	// --conn value: set num of UDP connections to server (default: 1)
	numconn := uint16(1)
	muxes := make([]struct {
		session mux.Mux
		ttl     time.Time
	}, numconn)

	for k := range muxes {
		muxes[k].session = waitConn()
		muxes[k].ttl = time.Now().Add(time.Duration(config.AutoExpireSeconds) * time.Second)
	}

	chScavenger := make(chan mux.Mux, 128)
	go scavenger(chScavenger, config.ScavengeTTL, lg)
	rr := uint16(0)
	for {
		p1, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		idx := rr % numconn

		// do auto expiration && reconnection
		if muxes[idx].session.IsClosed() || (config.AutoExpireSeconds > 0 && time.Now().After(muxes[idx].ttl)) {
			chScavenger <- muxes[idx].session
			muxes[idx].session = waitConn()
			muxes[idx].ttl = time.Now().Add(time.Duration(config.AutoExpireSeconds) * time.Second)
		}

		go handleClient(muxes[idx].session, p1, lg)
		rr++
	}
}

type scavengeSession struct {
	session mux.Mux
	ts      time.Time
}

func scavenger(ch chan mux.Mux, ttl int, lg logger.Logger) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	var sessionList []scavengeSession
	for {
		select {
		case sess := <-ch:
			sessionList = append(sessionList, scavengeSession{sess, time.Now()})
			lg.Infof("session marked as expired %s", sess.RemoteAddr())
		case <-ticker.C:
			var newList []scavengeSession
			for k := range sessionList {
				s := sessionList[k]
				if s.session.NumStreams() == 0 || s.session.IsClosed() {
					lg.Infof("session normally closed %s", s.session.RemoteAddr())
					_ = s.session.Close()
				} else if ttl >= 0 && time.Since(s.ts) >= time.Duration(ttl)*time.Second {
					lg.Infof("session reached scavenge ttl %s", s.session.RemoteAddr())
					_ = s.session.Close()
				} else {
					newList = append(newList, sessionList[k])
				}
			}
			sessionList = newList
		}
	}
}
