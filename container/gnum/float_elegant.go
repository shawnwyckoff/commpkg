package gnum

import (
	"math"
	"strconv"
)

// ElegantFloat is used to format float to better looking and enough precision string
// it supports NaN when json Marshal/Unmarshal

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
	var r []ElegantFloat
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
		switch string(b) {
		case "NaN":
			t.val = math.NaN()
			return nil
		case "+Inf":
			t.val = math.Inf(1)
			return nil
		case "-Inf":
			t.val = math.Inf(-1)
			return nil
		default:
			return err
		}
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
