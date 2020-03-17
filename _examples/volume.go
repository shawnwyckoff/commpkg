package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/dsa/volume"
)

func main() {
	vol, err := volume.ParseString("10 MB")
	fmt.Println(vol.String(), err)
}