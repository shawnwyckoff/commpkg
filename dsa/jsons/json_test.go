package jsons

import (
	"github.com/tidwall/gjson"
	"testing"
)

const (
	demo_json1 = `{"Time":"2018-04-10T16:14:08.364623+08:00","Content":123}`
	demo_json2 = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
)

func TestGet(t *testing.T) {
	jv := Get(demo_json1, "Content")
	if !jv.Exists() {
		t.Error("Get error, not exists")
		return
	}
	if jv.Type != gjson.Number {
		t.Error("Get error, type error")
		return
	}
}
