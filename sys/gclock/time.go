package gclock

import (
	"github.com/shawnwyckoff/gpkg/container/gstring"
	"time"
)

func Sub(t time.Time, d time.Duration) time.Time {
	return t.Add(0 - d)
}

func TimeToIntYYYYMMDDHHMM(t time.Time) int {
	return (t.Year() * 100000000) + (int(t.Month()) * 1000000) + (t.Day() * 10000) + (t.Hour() * 100) + t.Minute()
}

func StringTimeZone(tm time.Time, tz time.Location) string {
	return tm.In(&tz).String()
}

var (
	EpochBeginTime time.Time = EpochSecToTime(0)              // 1970-01-01 00:00:00 +0000 UTC
	EpochBeginDate Date      = Date(19700101)                 // 0001-01-01 00:00:00 +0000 UTC
	ZeroTime       time.Time = time.Time{}                    // 0001-01-01 00:00:00 +0000 UTC
	ZeroDate       Date      = TimeToDate(ZeroTime, time.UTC) // 0001-01-01 00:00:00 +0000 UTC
	ZeroYearMonth  YearMonth = 0                              // 0000-00
	ZeroDateRange  DateRange = DateRange{Begin: ZeroDate, End: ZeroDate}
)

func AfterEqual(a, b time.Time) bool {
	return a.After(b) || a.Equal(b)
}

func BeforeEqual(a, b time.Time) bool {
	return a.Before(b) || a.Equal(b)
}

func MinTime(a time.Time, b ...time.Time) time.Time {
	min := a
	for _, v := range b {
		if v.Before(min) {
			min = v
		}
	}
	return min
}

func MaxTime(a time.Time, b ...time.Time) time.Time {
	max := a
	for _, v := range b {
		if v.After(max) {
			max = v
		}
	}
	return max
}

func FormatDay(tm time.Time) string {
	return tm.Format("2006-01-02")
}

const TimeLayout_MM_DD_HH = "15:04:05"
const TimeLayout_YYYY = "2006"
const TimeLayout_YYYY_MM = "2006-01"
const TimeLayout_YYYY_MM_DD = "2006-01-02"
const TimeLayout_YYYY_MM_DD_HH = "2006-01-02 15"
const TimeLayout_YYYY_MM_DD_HH_mm = "2006-01-02 15:04"
const TimeLayout_YYYY_MM_DD_HH_mm_SS = "2006-01-02 15:04:05"
const TimeLayout_YYYY_MM_DD_HH_mm_SS_NS = "2006-01-02 15:04:05.999999999"
const TimeLayout_FULL = "2006-01-02 15:04:05.999999999 -0700 MST"

// ElegantTime is the time.Time with JSON marshal and unmarshal capability
type ElegantTime struct {
	val    time.Time
	layout string
}

func NewElegantTime(tm time.Time, layout string) ElegantTime {
	if layout == "" {
		layout = TimeLayout_FULL
	}
	return ElegantTime{val: tm, layout: layout}
}

func NewElegantTimeArray(tms []time.Time, layout string) []ElegantTime {
	var r []ElegantTime
	for _, v := range tms {
		r = append(r, NewElegantTime(v, layout))
	}
	return r
}

func (t *ElegantTime) Raw() time.Time {
	return t.val
}

func (t *ElegantTime) SetLayout(layout string) {
	t.layout = layout
}

// UnmarshalJSON will unmarshal using 2006-01-02T15:04:05+07:00 layout
func (t *ElegantTime) UnmarshalJSON(b []byte) error {
	val, err := ParseDatetimeStringFuzz(string(b))
	if err != nil {
		return err
	}

	t.val = val
	t.layout = TimeLayout_FULL
	return nil
}

// MarshalJSON will marshal using 2006-01-02T15:04:05+07:00 layout
func (t *ElegantTime) MarshalJSON() ([]byte, error) {
	return t.JSON()
}

func (t *ElegantTime) JSONAutoDetect() ([]byte, error) {
	s := t.val.Format(t.DetectBestLayout())
	return []byte(`"` + s + `"`), nil
}

func (t *ElegantTime) JSON() ([]byte, error) {
	if t.layout == "" {
		t.layout = TimeLayout_FULL
	}
	s := t.val.Format(t.layout)
	return []byte(`"` + s + `"`), nil
}

func (t *ElegantTime) DetectBestLayout() string {
	hasMonth := t.val.Month() != 1
	hasDay := t.val.Day() != 1
	hasHour := t.val.Hour() != 0
	hasMinute := t.val.Minute() != 0
	hasSecond := t.val.Second() != 0
	hasNanosecond := t.val.Nanosecond() != 0
	hasNonUTCTimeZone := t.val.Location() != time.UTC

	// count start with month
	cswNanosecond := 0
	if hasNanosecond {
		cswNanosecond++
	}

	cswSecond := 0
	if hasSecond {
		cswSecond++
	}
	cswSecond += cswNanosecond

	cswMinute := 0
	if hasMinute {
		cswMinute++
	}
	cswMinute += cswSecond

	cswHour := 0
	if hasHour {
		cswHour++
	}
	cswHour += cswMinute

	cswDay := 0
	if hasDay {
		cswDay++
	}
	cswDay += cswHour

	cswMonth := 0
	if hasMonth {
		cswMonth++
	}
	cswMonth += cswDay

	layoutTZ := ""
	if hasNonUTCTimeZone {
		layoutTZ = " -0700 MST"
	}
	if !hasMonth && !hasDay && !hasHour && !hasMinute && !hasSecond && !hasNanosecond {
		return TimeLayout_YYYY + layoutTZ
	} else if !hasDay && !hasHour && !hasMinute && !hasSecond && !hasNanosecond {
		return TimeLayout_YYYY_MM + layoutTZ
	} else if !hasHour && !hasMinute && !hasSecond && !hasNanosecond {
		return TimeLayout_YYYY_MM_DD + layoutTZ
	} else if !hasMinute && !hasSecond && !hasNanosecond {
		return TimeLayout_YYYY_MM_DD_HH + layoutTZ
	} else if !hasSecond && !hasNanosecond {
		return TimeLayout_YYYY_MM_DD_HH_mm + layoutTZ
	} else if !hasNanosecond {
		return TimeLayout_YYYY_MM_DD_HH_mm_SS + layoutTZ
	} else {
		return TimeLayout_YYYY_MM_DD_HH_mm_SS_NS + layoutTZ
	}
}

func DetectBestLayout(in []ElegantTime) string {
	if in == nil || len(in) == 0 {
		return ""
	}

	_LAYOUT_TIMEZONE_ := " -0700 MST"

	layoutHead := ""
	layoutTimeZone := ""
	for _, v := range in {
		tmp := v.DetectBestLayout()
		tmpHead := tmp
		tmpTZ := ""
		if gstring.EndWith(tmp, _LAYOUT_TIMEZONE_) {
			tmpHead = gstring.RemoveTail(tmp, len(_LAYOUT_TIMEZONE_))
			tmpTZ = _LAYOUT_TIMEZONE_
		}

		if len(tmpHead) > len(layoutHead) {
			layoutHead = tmpHead
		}
		if len(tmpTZ) > 0 {
			layoutTimeZone = _LAYOUT_TIMEZONE_
		}
	}
	return layoutHead + layoutTimeZone
}

func DetectBestLayoutRaw(in []time.Time) string {
	inET := []ElegantTime{}
	for _, v := range in {
		inET = append(inET, NewElegantTime(v, ""))
	}
	return DetectBestLayout(inET)
}
