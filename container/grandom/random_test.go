package grandom

import "testing"

func TestRandomInt(t *testing.T) {
	RandomInt(0, 0)
}

func TestRandomString(t *testing.T) {
	s := RandomString(9)
	if len(s) != 9 {
		t.Errorf("RandomString error")
		return
	}
}
