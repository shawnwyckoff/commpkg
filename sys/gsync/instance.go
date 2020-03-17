package gsync

import (
	"github.com/marcsauter/single"
	"github.com/pkg/errors"
)

type SingleInstanceLock struct {
	lock *single.Single
}

func NewSingleInstanceLock(appName string) *SingleInstanceLock {
	res := SingleInstanceLock{}
	res.lock = single.New(appName)
	return &res
}

func (l *SingleInstanceLock) IsUnique() (bool, error) {
	if err := l.lock.CheckLock(); err != nil && err == single.ErrAlreadyRunning {
		return false, nil
	} else if err != nil {
		// Another error occurred, might be worth handling it as well
		return false, errors.Errorf("failed to acquire exclusive app lock: %v", err)
	}

	return true, nil
}

func (l *SingleInstanceLock) UnLock() error {
	return l.lock.TryUnlock()
}
