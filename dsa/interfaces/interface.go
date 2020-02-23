package interfaces

import (
	"log"
	"reflect"
	"time"
)

type CommonFunc struct{}

var cf CommonFunc

func (c *CommonFunc) Merge2(s ...[]interface{}) (result []interface{}) {
	switch len(s) {
	case 0:
		break
	case 1:
		result = s[0]
		break
	default:
		s1 := s[0]
		s2 := cf.Merge2(s[1:]...) //...将数组元素打散
		result = make([]interface{}, len(s1)+len(s2))
		copy(result, s1)
		copy(result[len(s1):], s2)
		break
	}

	return
}

func Merge(s ...[]interface{}) (result []interface{}) {
	switch len(s) {
	case 0:
		break
	case 1:
		result = s[0]
		break
	default:
		s1 := s[0]
		s2 := Merge(s[1:]...) // ...可以将数组元素打散
		result = make([]interface{}, len(s1)+len(s2))
		copy(result, s1)
		copy(result[len(s1):], s2)
		break
	}

	return result
}

/**
  @retry  重试次数
  @method 调用的函数，比如: api.GetTicker ,注意：不是api.GetTicker(...)
  @params 参数,顺序一定要按照实际调用函数入参顺序一样
  @return 返回
*/
func ReCallItfc(retry int, method interface{}, params ...interface{}) interface{} {

	invokeM := reflect.ValueOf(method)
	if invokeM.Kind() != reflect.Func {
		panic("method not a function")
		return nil
	}

	var value []reflect.Value = make([]reflect.Value, len(params))
	var i int = 0
	for ; i < len(params); i++ {
		value[i] = reflect.ValueOf(params[i])
	}

	var retV interface{}
	var retryC int = 0
_CALL:
	if retryC > 0 {
		log.Println("sleep....", time.Duration(retryC*200*int(time.Millisecond)))
		time.Sleep(time.Duration(retryC * 200 * int(time.Millisecond)))
	}

	retValues := invokeM.Call(value)

	for _, vl := range retValues {
		if vl.Type().String() == "error" {
			if !vl.IsNil() {
				log.Println(vl)
				retryC++
				if retryC <= retry {
					log.Printf("Invoke Method(%s) Error , Begin Retry Call [%d] ...", invokeM.String(), retryC)
					goto _CALL
				} else {
					panic("Invoke Method Fail ???" + invokeM.String())
				}
			}
		} else {
			retV = vl.Interface()
		}
	}

	return retV
}
