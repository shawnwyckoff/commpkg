package gclock

import (
	"github.com/pkg/errors"
	"time"
)

type Clock interface {
	Name() string
	Now() time.Time
	Sleep(d time.Duration)
}

func NewClock(clockName string) (Clock, error) {
	switch clockName {
	case "system":
		return GetSysClock(), nil
	case "ntp":
		return GetNtpClockONLINE()
	case "mock":
		return NewMockClock(time.Now(), time.UTC), nil
	}
	return nil, errors.Errorf("unsupported clock name %s", clockName)
}
