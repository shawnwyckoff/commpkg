package gerror

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/container/gstring"
	"runtime"
	"strconv"
	"strings"
)

type (
	GError struct {
	}
)

var (
	ErrNil = error(nil)
	ErrNotFound = errors.Errorf("not found") // this is not a really run error, it means Database/Collection not exist in mongodb.
)

// github.com/txgruppi/werr seems a better error implement

// https://github.com/aletheia7/errors/blob/master/e.go

// support "github/pkg/errors"
// not support standard "errors"
// 有时，即使使用了pkg/errors, 用errors.New得到的err也无法获取堆栈（有时可以），但是用errors.Errorf可以
func getStack(err error) string {
	//return probe.Trace(err).Error() + "heheda"
	//err.(*errors.Error).ErrorStack()
	if err == nil {
		return ""
	}
	return gstring.RemoveFirstLines(fmt.Sprintf("%+v\n", err), 1)
}

// https://github.com/fd0/probe 这个包可以跟踪错误堆栈？
// xqtp.listen("")的时候就无法打印堆栈

func New(format string, a ...interface{}) error {
	return errors.Errorf(format, a...)
}

func GetStack(err error) string {
	stack := getStack(err)
	if len(stack) == 0 || len(strings.Replace(stack, " ", "", -1)) == 0 {
		// Get call stack.
		for i := 2; i < 10; i++ {
			pc, file, line, _ := runtime.Caller(i) // 调用文件名和行号
			if len(file) == 0 {
				break
			}
			if gstring.EndWith(file, "runtime.main") {
				break
			}
			f := runtime.FuncForPC(pc) //调用包名.函数名
			stack = stack + "    -> " + file + ":" + strconv.FormatInt(int64(line), 10) + " " + f.Name()
		}
	}
	return stack
}

func Error2Fatal(err error) error {
	return errors.New("fatal:" + err.Error())
}

func Errorf(format string, arguments ...interface{}) error {
	return errors.Errorf(format, arguments...)
}

func Fatalf(format string, arguments ...interface{}) error {
	return errors.Errorf("fatal:"+format, arguments...)
}

func IsFatal(err error) bool {
	return gstring.StartWith(err.Error(), "fatal:")
}
