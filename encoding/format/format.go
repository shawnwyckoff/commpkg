package format

import (
	"fmt"
	"github.com/floyernick/fleep-go"
	"io/ioutil"
	"strings"
)

func GetFormat(path string) (string, error) {
	file, _ := ioutil.ReadFile(path) // Reads PNG file
	info, _ := fleep.GetInfo(file)   // Gets file format
	fmt.Println(info.Type)
	return strings.Join(info.Type, ","), nil
}
