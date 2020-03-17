package gstring

import (
	"math/rand"
	"strings"
	"time"
)

func ArrayEqual(a, b []string, orderSensitive bool) bool {
	return false
}

func FindMaxLength(elements []string) []string {
	if len(elements) <= 1 {
		return elements
	}

	maxlen := 0
	for _, str := range elements {
		if len(str) > maxlen {
			maxlen = len(str)
		}
	}

	rst := make([]string, 0)
	for _, val := range elements {
		if maxlen == len(val) {
			rst = append(rst, val)
		}
	}
	return rst
}

func RemoveSpaces(elements []string) []string {
	r := []string{}
	for _, v := range elements {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

func RemoveDuplicate(elements []string) []string {
	if len(elements) <= 1 {
		return elements
	}

	// another way to initialize map
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	var result []string
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func ToLower(elements []string) []string {
	if len(elements) == 0 {
		return []string{}
	}

	rst := []string{}
	for _, v := range elements {
		rst = append(rst, strings.ToLower(v))
	}
	return rst
}

func ToUpper(elements []string) []string {
	if len(elements) == 0 {
		return []string{}
	}

	rst := []string{}
	for _, v := range elements {
		rst = append(rst, strings.ToUpper(v))
	}
	return rst
}

func RemoveByValue(elements []string, toRemove string) []string {
	if len(elements) == 0 {
		return nil
	}

	result := make([]string, 0)
	for _, val := range elements {
		if toRemove != val {
			result = append(result, val)
		}
	}
	return result
}

func RemoveByValues(elements, toRemove []string) []string {
	if len(elements) == 0 {
		return nil
	}
	if len(toRemove) == 0 {
		return elements
	}

	result := make([]string, 0)
	for _, val := range elements {
		if CountByValue(toRemove, val) <= 0 {
			result = append(result, val)
		}
	}
	return result
}

func CountByValue(elements []string, toFind string) int {
	var result int = 0
	for _, val := range elements {
		if toFind == val {
			result++
		}
	}
	return result
}

func Contains(elements []string, toFind string) bool {
	for _, val := range elements {
		if toFind == val {
			return true
		}
	}
	return false
}

// Random sort
func Shuffle(elements []string) []string {
	final := make([]string, len(elements))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(elements))

	for i, v := range perm {
		final[v] = elements[i]
	}

	return final
}

func Merge(elements1, elements2 []string) []string {
	if elements1 == nil {
		return elements2
	}
	if elements2 == nil {
		return elements1
	}
	result := elements1
	for i := range elements2 {
		result = append(result, elements2[i])
	}
	return result
}

// 插入到前面
func Prepend(slice []string, elems ...string) []string {
	slice = append(elems, slice...)
	return slice
}

func Equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func IndexInArray(ss []string, tofind string) int {
	for i := 0; i < len(ss); i++ {
		if ss[i] == tofind {
			return i
		}
	}
	return -1
}
