package p2pdns

import "sync"

type Query struct {
	address []string
	mu      sync.RWMutex
}

func Join(domain string, addressWithPort string) error {

}

func Lookup(domain string) (*Query, error) {

}

func (s *Query) GetAddress() []string {

}
