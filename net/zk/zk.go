package zk

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

type (
	ZK struct {
		conn *zk.Conn
		chEvt <-chan zk.Event
	}
)

func Dial(servers []string, timeout time.Duration) (*ZK, error) {
	conn, chEvt, err := zk.Connect(servers, timeout)
	if err != nil {
		return nil, err
	}
	return &ZK{conn:conn, chEvt:chEvt}, nil
}

func (zk *ZK) RegisterService(name, addr string) error {
	return nil
}

func (zk *ZK) QueryService(name string) ([]string, error) {
	return nil, nil
}

