package gvolume

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gopkg/container/gspeed"
	"github.com/shawnwyckoff/gopkg/container/gstring"
	"strings"
)

// Map between speed unit of byte and bits size
const (
	KB Volume = Volume(gspeed.KB)
	MB        = Volume(gspeed.MB)
	GB        = Volume(gspeed.GB)
	TB        = Volume(gspeed.TB)
	PB        = Volume(gspeed.PB)
	EB        = Volume(gspeed.EB)
	ZB        = Volume(gspeed.ZB)
	YB        = Volume(gspeed.YB)
)

type Volume gspeed.Speed

func FromByteSize(size float64) (Volume, error) {
	v, err := gspeed.FromBytes(size)
	return (Volume)(v), err
}

func (v *Volume) BiggerThan(v2 Volume) bool {
	return (*gspeed.Speed)(v).BiggerThan((gspeed.Speed)(v2))
}

func (v *Volume) SmallerThan(v2 Volume) bool {
	return (*gspeed.Speed)(v).SmallerThan((gspeed.Speed)(v2))
}

func (v *Volume) Equals(v2 Volume) bool {
	return (*gspeed.Speed)(v).Equals((gspeed.Speed)(v2))
}

func (v Volume) String() string {
	return ((*gspeed.Speed)(&v)).StringWithByteUnit()
}

func (v Volume) Bits() float64 {
	return float64(v)
}

func (v Volume) Bytes() uint64 {
	return uint64(v) / 8
}

func (v Volume) MBytes() float64 {
	return float64(v / MB)
}

// volume string sample: "2M" "2MB" "2Mbyte" "2Mbytes"
func ParseString(volume string) (Volume, error) {
	// Returns error if per second found
	t := strings.ToLower(volume)
	t = strings.TrimSpace(t)
	if gstring.EndWith(t, "/s") || gstring.EndWith(t, "ps") {
		return Volume(0), errors.New("V \"" + volume + "\" syntax error")
	}

	// Returns error if b/bit found
	t = strings.TrimSpace(volume)
	t = strings.ToLower(t)
	if strings.Contains(t, "bit") || strings.Contains(t, "bits") || (!strings.Contains(t, "byte") && strings.Contains(volume, "b")) {
		return Volume(0), errors.New("V \"" + volume + "\" syntax error")
	}

	// Parse with speed library
	v, err := gspeed.ParseString(volume)
	return (Volume)(v), err
}

func (v Volume) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", v.String())), nil
}

func (v *Volume) UnmarshalJSON(b []byte) error {
	str := string(b)
	if len(str) <= 1 {
		return errors.Errorf("invalid json volume '%s'", v)
	}
	if str[0] != '"' || str[len(str)-1] != '"' {
		return errors.Errorf("invalid json volume '%s'", v)
	}
	str = gstring.RemoveHead(str, 1)
	str = gstring.RemoveTail(str, 1)
	speed, err := ParseString(str)
	if err != nil {
		return err
	}
	*v = speed
	return nil
}
