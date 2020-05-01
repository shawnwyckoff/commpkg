package gnum

import (
	"math"
	"testing"
)

func TestElegantFloat_JSON(t *testing.T) {
	var f1 = NewElegantFloat(54321.00153456789123, -1)
	f1.SetHumanReadPrec(2)
	b, err := f1.JSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "54321.0015" || f1.String() != "54321.0015" {
		t.Error("f1 format error, should be 54321.0015")
	}

	var f2 = NewElegantFloat(0.001253456789123, -1)
	f2.SetHumanReadPrec(2)
	b, err = f2.JSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "0.0013" || f2.String() != "0.0013" {
		t.Error("f2 format error, should be 0.0013")
	}

	var f3 = NewElegantFloat(math.NaN(), -1)
	b, err = f3.JSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "NaN" {
		t.Error("f3 format error, should be NaN")
	}

	var f4 = NewElegantFloat(math.Inf(1), -1)
	b, err = f4.JSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "+Inf" {
		t.Error("f4 format error, should be +Inf")
	}

	var f5 = NewElegantFloat(math.Inf(-1), -1)
	b, err = f5.JSON()
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "-Inf" {
		t.Error("f5 format error, should be -Inf")
	}
}
