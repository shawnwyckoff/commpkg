package interfaces

import (
	"reflect"
)

func Parse(x interface{}) (typeName string, isPointer bool) {
	if t := reflect.TypeOf(x); t.Kind() == reflect.Ptr {
		return t.Elem().Name(), true
	} else {
		return t.Name(), false
	}
}

func Type(x interface{}) string {
	if x == nil {
		return "nil"
	}
	return reflect.TypeOf(x).String()
	/*
		if t := reflect.TypeOf(x); t.Kind() == reflect.Ptr {
			return "*" + t.Elem().Name()
		} else {
			return t.Name()
		}*/
}
