package main

import (
	"fmt"
	"github.com/shawnwyckoff/commpkg/apputil/logz"
	"github.com/shawnwyckoff/commpkg/dsa/stringz"
	"github.com/shawnwyckoff/commpkg/sys/fs"
	"os"
	"path/filepath"
	"strings"
)

func main()  {
	goPath := os.Getenv("GOPATH")
	logz.AssertTrue(goPath != "")

	repoPath := "github.com/shawnwyckoff/commpkg"
	commpkgPath := filepath.Join(goPath, "src", repoPath)
	dirs, _, err := fs.ListDir(commpkgPath)
	logz.AssertOk(err, "ListDir")

	res := map[string][]string{}
	for _, v := range dirs {
		ss := strings.Split(v, repoPath)
		if len(ss) != 2 {
			continue
		}
		subItems := strings.Split(ss[1], "/")
		subItems = stringz.RemoveByValue(subItems, "")
		if len(subItems) == 2 {
			original := res[subItems[0]]
			original = append(original, subItems[1])
			res[subItems[0]] = original
		}
	}

	pkgCount := 0
	for pkgSort, pkgs := range res {
		if pkgSort == ".git" {
			continue
		}
		fmt.Println("\n**" + pkgSort + "**\n")
		for _, pkg := range pkgs {
			pkgCount++
			fmt.Print(pkg + "  ")
		}
		fmt.Println("")
	}
	fmt.Println("\ntotal:", pkgCount)
}
