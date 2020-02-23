package main

import (
	"fmt"
	"github.com/shawnwyckoff/commpkg/dsa/volume"
)

func main() {
	vol, err := volume.ParseString("10 MB")
	fmt.Println(vol.String(), err)
}