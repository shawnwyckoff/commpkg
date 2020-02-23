package interfaces

import (
	serr "errors"
	"github.com/pkg/errors"
	"testing"
)

func TestTypeString(t *testing.T) {
	src := "123"
	typeString := Type(src)
	if typeString != "string" {
		t.Errorf("Parse string error, returns %s", typeString)
	}

	typeString = Type(&src)
	if typeString != "*string" {
		t.Errorf("Parse *string error, returns %s", typeString)
	}

	num := 123.456
	typeString = Type(num)
	if typeString != "float64" {
		t.Errorf("Parse float64 error, returns %s", typeString)
	}

	err := serr.New("this is a standard error")
	typeString = Type(err)
	if typeString != "*errors.errorString" {
		t.Errorf("Parse standard error type error, returns %s", typeString)
	}

	err = errors.Errorf("this is a extended error")
	typeString = Type(err)
	if typeString != "*errors.fundamental" {
		t.Errorf("Parse extended error type error, returns %s", typeString)
	}

	err = errors.New("this is a extended error too")
	typeString = Type(err)
	if typeString != "*errors.fundamental" {
		t.Errorf("Parse extended error too type error, returns %s", typeString)
	}

	type myStruct struct{}
	ms := myStruct{}
	typeString = Type(&ms)
	if typeString != "*interfaces.myStruct" {
		t.Errorf("Parse *myStruct type error, returns %s", typeString)
	}
}
