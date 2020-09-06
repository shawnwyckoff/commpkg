package gjson

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/tidwall/gjson"
	"time"
)

type JsonValue gjson.Result

func Get(json, path string) JsonValue {
	return JsonValue(gjson.Get(json, path))
}

func Set(jsonstr string, path []string, val interface{}) (string, error) {
	js, err := simplejson.NewJson([]byte(jsonstr))
	if err != nil {
		return "", err
	}
	js.SetPath(path, val)
	b, err := js.MarshalJSON()
	return string(b), err
}

func (v JsonValue) Exists() bool {
	return gjson.Result(v).Exists()
}

func (v JsonValue) String() string {
	return gjson.Result(v).String()
}

func (v JsonValue) Time() time.Time {
	return gjson.Result(v).Time()
}

func (v JsonValue) Bool() bool {
	return gjson.Result(v).Bool()
}

func (v JsonValue) Float() float64 {
	return gjson.Result(v).Float()
}

func (v JsonValue) Int64() int64 {
	return int64(gjson.Result(v).Float())
}

func MarshalAndPrintln(x interface{}) {
	buf, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}
