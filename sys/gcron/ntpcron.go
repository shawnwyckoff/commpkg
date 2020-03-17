package gcron

import (
	"github.com/shawnwyckoff/gpkg/sys/gclock"
	"time"
)

type NtpCron struct {
	nc          *gclock.NtpClock
	triggerTime time.Time
}

func NewNtpCronONLINE(ntpClock *gclock.NtpClock, triggerTime time.Time) (*NtpCron, error) {
	nc, err := gclock.GetNtpClockONLINE()
	if err != nil {
		return nil, err
	}
	return &NtpCron{nc: nc, triggerTime: triggerTime}, nil
}

func NewNtpCronWithClock(ntpClock *gclock.NtpClock, triggerTime time.Time) (*NtpCron, error) {
	return &NtpCron{nc: ntpClock, triggerTime: triggerTime}, nil
}

func (sc *NtpCron) Wait() {
	for {
		sleepMillis := 100 // default sleep 5000 milliseconds for each loop
		dur := sc.triggerTime.Sub(sc.nc.Now())
		durMillis := gclock.NsecToMillis(dur.Nanoseconds())
		if durMillis < 40 {
			break
		}
		time.Sleep(gclock.MillisToDuration(int64(sleepMillis)))
	}
}
