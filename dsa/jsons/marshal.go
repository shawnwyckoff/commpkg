package jsons

import (
	"encoding/json"
	"fmt"
	"github.com/mailru/easyjson" //
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// github.com/mailru/easyjson is much faster than encoding/json
// github.com/pquerna/ffjson 2x-3x faster than encoding/json

func Marshal(v interface{}, indent bool, errfmt string) []byte {
	var buf []byte
	err := error(nil)
	if indent {
		buf, err = json.MarshalIndent(v, "", "\t")
	} else {
		buf, err = json.Marshal(v)
	}
	if err != nil {
		return []byte(fmt.Sprintf(errfmt, err.Error()))
	}
	return buf
}

func MarshalString(v interface{}, indent bool, errfmt string) string {
	return string(Marshal(v, indent, errfmt))
}

func MarshalStringDefault(v interface{}, indent bool) string {
	return MarshalString(v, indent, `{"Error":"%s"}`)
}

func MarshalIndent(v interface{}) (string, error) {
	buf, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func MarshalFast(v easyjson.Marshaler) ([]byte, error) {
	return easyjson.Marshal(v)
}

// JSONEncode encodes structure data into JSON
func JSONEncode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// JSONDecode decodes JSON data into a structure
func JSONDecode(data []byte, to interface{}) error {
	if !strings.Contains(reflect.ValueOf(to).Type().String(), "*") {
		return errors.New("json decode error - memory address not supplied")
	}
	return json.Unmarshal(data, to)
}
