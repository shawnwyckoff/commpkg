package gnum

import (
	"encoding/binary"
	"math"
	"strconv"
	"strings"
)

var (
	PosInf = math.Inf(1)
	NegInf = math.Inf(-1)
)

const float64EqualityThreshold = 1e-9

func FloatAlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

// Converts the float64 into an uint64 without changing the bits, it's the way the bits are interpreted that change.
// big endian
// references:
// https://stackoverflow.com/questions/37758267/golang-float64bits
// https://stackoverflow.com/questions/43693360/convert-float64-to-byte-array
func Float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

// little endian
func Float64ToByteLE(f float64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

const (
	invalidPrec = -2
	defaultPrec = -1
)

// ElegantFloat is used to format float to better looking and enough precision string
//
// original            -> prec: 2       -> humanReadPrec: 2
// 37                  -> 37.00         -> 37.00
// 12237.89374         -> 12237.89      -> 12237.89
// 3.3483300000000003  -> 3.35          -> 3.35
// 0.00883300000000003 -> 0.01          -> 0.0088
// 0.000012800003      -> 0.00          -> 0.000013
type ElegantFloat struct {
	val  float64
	prec int
}

func NewElegantFloat(val float64, prec int) ElegantFloat {
	return ElegantFloat{val: val, prec: prec}
}

func NewElegantFloatArray(vals []float64, prec int) []ElegantFloat {
	r := []ElegantFloat{}
	for _, v := range vals {
		r = append(r, NewElegantFloat(v, prec))
	}
	return r
}

func NewElegantFloatPtrArray(vals []*float64, prec int) []*ElegantFloat {
	var r []*ElegantFloat
	for _, v := range vals {
		if v == nil {
			r = append(r, nil)
		} else {
			nf := NewElegantFloat(*v, prec)
			r = append(r, &nf)
		}
	}
	return r
}

func NewElegantFloatPtrArray2(vals []float64, prec int) []*ElegantFloat {
	var r []*ElegantFloat
	for _, v := range vals {
		nf := NewElegantFloat(v, prec)
		r = append(r, &nf)
	}
	return r
}

func NewElegantFloatPtrArray3(vals []float64, prec int, nilVal float64) []*ElegantFloat {
	var r []*ElegantFloat
	for _, v := range vals {
		if v == nilVal {
			r = append(r, nil)
		} else {
			nf := NewElegantFloat(v, prec)
			r = append(r, &nf)
		}
	}
	return r
}

func ElegantFloatArrayToFloatArray(in []ElegantFloat) []float64 {
	var r []float64
	for i := range in {
		r = append(r, in[i].val)
	}
	return r
}

func (t *ElegantFloat) SetPrec(prec int) {
	if prec > invalidPrec {
		t.prec = prec
	}
}

func (t *ElegantFloat) SetHumanReadPrec(humanReadPrec int) {
	t.prec = DetectPrecByHumanReadPrec(t.val, humanReadPrec)
}

func (t *ElegantFloat) Raw() float64 {
	return t.val
}

// UnmarshalJSON will unmarshal using 2006-01-02T15:04:05+07:00 layout
func (t *ElegantFloat) UnmarshalJSON(b []byte) error {
	val, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		return err
	}

	t.val = val
	t.prec = invalidPrec
	return nil
}

// MarshalJSON will marshal using 2006-01-02T15:04:05+07:00 layout
func (t ElegantFloat) MarshalJSON() ([]byte, error) {
	return t.JSON()
}

func (t *ElegantFloat) JSON() ([]byte, error) {
	if t.prec <= invalidPrec {
		t.prec = defaultPrec
	}
	s := strconv.FormatFloat(t.val, 'f', t.prec, 64)
	return []byte(s), nil
}

func (t *ElegantFloat) String() string {
	return strconv.FormatFloat(t.val, 'f', t.prec, 64)
}

func DetectPrecByHumanReadPrec(val float64, humanReadPrec int) int {
	if humanReadPrec <= invalidPrec {
		return defaultPrec
	}

	// example t.val is 2.001263645807, humanReadPrec is 3
	s := strconv.FormatFloat(val, 'f', -1, 64)
	if strings.Index(s, ".") > 0 {
		s = strings.Split(s, ".")[1] // s is "001263645807"
		for i := range s {
			if s[i] != '0' { // i is 2
				return i + humanReadPrec // returns 5
			}
		}
		return humanReadPrec
	} else {
		return humanReadPrec
	}
}

func DetectMaxPrec(vals []ElegantFloat, humanReadPrec int) int {
	r := defaultPrec
	tmp := defaultPrec
	for _, v := range vals {
		tmp = DetectPrecByHumanReadPrec(v.val, humanReadPrec)
		if tmp > r {
			r = tmp
		}
	}
	return r
}

func DetectMaxPrecRaw(vals []float64, humanReadPrec int) int {
	r := defaultPrec
	tmp := defaultPrec
	for _, v := range vals {
		tmp = DetectPrecByHumanReadPrec(v, humanReadPrec)
		if tmp > r {
			r = tmp
		}
	}
	return r
}
