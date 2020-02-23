package sysinfo

import "github.com/pkg/errors"

func GetSystemProxy() (string, bool, error) {
	// TODO
	return "", false, errors.Errorf("unsupported for now")
}

func SetSystemProxy(defaultServer string, enabled bool) error {
	// TODO
	return errors.Errorf("unsupported for now")
}
