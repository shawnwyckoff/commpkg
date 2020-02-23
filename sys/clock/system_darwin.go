package clock

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/commpkg/sys/users"
	"syscall"
	"time"
)

func SetSystemTimeROOT(t time.Time) error {
	var tv syscall.Timeval
	tv.Sec = t.Unix()
	tv.Usec = 0
	if err := syscall.Settimeofday(&tv); err != nil {
		isAdmin, err2 := users.IsRunAsAdmin()
		if err2 == nil && !isAdmin {
			return errors.Errorf(err.Error() + ", modifying system time requires administrator privileges")
		} else {
			return err
		}
	}
	return nil
}
