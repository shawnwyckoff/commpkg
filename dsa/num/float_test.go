package num

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/dsa/jsons"
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
}

func TestFloatAlmostEqual(t *testing.T) {
	type S struct {
		A float64
	}
	s := S{A: 2.345}
	fmt.Println(jsons.MarshalStringDefault(s, true))
	fmt.Println(fmt.Sprintf("%.5f", s.A))
}

func TestDetectMaxPrec(t *testing.T) {
	fmt.Println(math.IsNaN(PosInf))
	fmt.Println(fmt.Sprintf("%f", NegInf))
}
