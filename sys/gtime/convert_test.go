package gtime

import (
	"github.com/shawnwyckoff/gopkg/apputil/gtest"
	"testing"
	"time"
)

func TestEpochSecToTime(t *testing.T) {
	tm := EpochSecToTime(0)
	if !IsEpochBeginning(tm) {
		t.Error("EpochSecToTime(0) returns sec", tm.Second(), "nsec", tm.Nanosecond())
		return
	}
}

func TestUnixNanoToTime(t *testing.T) {
	now := time.Now()
	nowUN := now.UnixNano()
	if !UnixNanoToTime(nowUN).Equal(now) {
		gtest.PrintlnExit(t, "UnixNanoToTime test error1")
	}
}
