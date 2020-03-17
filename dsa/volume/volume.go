package volume

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/dsa/speed"
	"github.com/shawnwyckoff/gpkg/dsa/stringz"
	"strings"
)

// Map between speed unit of byte and bits size
const (
	KB Volume = Volume(speed.KB)
	MB        = Volume(speed.MB)
	GB        = Volume(speed.GB)
	TB        = Volume(speed.TB)
	PB        = Volume(speed.PB)
	EB        = Volume(speed.EB)
	ZB        = Volume(speed.ZB)
	YB        = Volume(speed.YB)
)

type Volume speed.Speed

func FromByteSize(size float64) (Volume, error) {
	v, err := speed.FromBytes(size)
	return (Volume)(v), err
}

func (v *Volume) BiggerThan(v2 Volume) bool {
	return (*speed.Speed)(v).BiggerThan((speed.Speed)(v2))
}

func (v *Volume) SmallerThan(v2 Volume) bool {
	return (*speed.Speed)(v).SmallerThan((speed.Speed)(v2))
}

func (v *Volume) Equals(v2 Volume) bool {
	return (*speed.Speed)(v).Equals((speed.Speed)(v2))
}

func (v Volume) String() string {
	return ((*speed.Speed)(&v)).StringWithByteUnit()
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
	if stringz.EndWith(t, "/s") || stringz.EndWith(t, "ps") {
		return Volume(0), errors.New("Volume \"" + volume + "\" syntax error")
	}

	// Returns error if b/bit found
	t = strings.TrimSpace(volume)
	t = strings.ToLower(t)
	if strings.Contains(t, "bit") || strings.Contains(t, "bits") || (!strings.Contains(t, "byte") && strings.Contains(volume, "b")) {
		return Volume(0), errors.New("Volume \"" + volume + "\" syntax error")
	}

	// Parse with speed library
	v, err := speed.ParseString(volume)
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
	str = stringz.RemoveHead(str, 1)
	str = stringz.RemoveTail(str, 1)
	speed, err := ParseString(str)
	if err != nil {
		return err
	}
	*v = speed
	return nil
}
