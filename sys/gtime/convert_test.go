package gtime

import (
	"testing"
)

func TestEpochSecToTime(t *testing.T) {
	tm := EpochSecToTime(0)
	if !IsEpochBeginning(tm) {
		t.Error("EpochSecToTime(0) returns sec", tm.Second(), "nsec", tm.Nanosecond())
		return
	}
}
