package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/dsa/gvolume"
)

func main() {
	vol, err := gvolume.ParseString("10 MB")
	fmt.Println(vol.String(), err)
}