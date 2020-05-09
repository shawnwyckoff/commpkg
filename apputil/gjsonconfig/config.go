package gjsonconfig

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gopkg/container/ginterface"
	"github.com/shawnwyckoff/gopkg/container/gstring"
	"github.com/shawnwyckoff/gopkg/sys/gfs"
	"github.com/shawnwyckoff/gopkg/sys/gproc"
	"io/ioutil"
	"runtime"
	"strings"
)

// Same Dir Same Name
func DefaultConfig() (string, error) {
	surffix := ".json"
	fn, err := gproc.SelfPath()
	if err != nil {
		return "", err
	}
	fnLower := strings.ToLower(fn)
	if (runtime.GOOS == "windows" && gstring.EndWith(fnLower, ".exe")) ||
		(runtime.GOOS == "darwin" && gstring.EndWith(fnLower, ".app")) ||
		(runtime.GOOS == "linux" && gstring.EndWith(fnLower, ".bin")) {
		fn = fn[0 : len(fn)-4]
	}
	return fn + surffix, nil
}

// Read same dir same name .json config file and unmarshal to a struct
// v is a pointer to structure
func DefaultUnmarshal(v interface{}) error {
	filename, err := DefaultConfig()
	if err != nil {
		return err
	}
	return Unmarshal(filename, v)
}

// Read same dir json conf file and unmarshal to a struct
// shortfn example: "*.json"
func Unmarshal(filename string, v interface{}) error {
	typeName, isPtr := ginterface.Parse(v)
	if !isPtr {
		return errors.Errorf("pointer needed for json.Unmarshal, but type %s is NOT a pointer", typeName)
	}

	pi, err := gfs.GetPathInfo(filename)
	if err != nil {
		return err
	}
	if pi.IsFolder || !pi.Exist {
		return errors.Errorf("config file '%s' is folder or not exist", filename)
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.Errorf("config file '%s' content is empty", filename)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
