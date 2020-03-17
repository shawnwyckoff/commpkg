// +build !windows,!linux

package gmtu

func check(addr string, size int) (bool, int, error) {
	return false, 0, Unimplemented
}
