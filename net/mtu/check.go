// +build !windows,!linux

package mtu

func check(addr string, size int) (bool, int, error) {
	return false, 0, Unimplemented
}
