package fs

import (
	"fmt"
	"os"
	"testing"
)

func TestMakeDir(t *testing.T) {
	if err := MakeDir("abc"); err != nil {
		t.Error(err)
		return
	}
	if err := RemoveDir("abc"); err != nil {
		t.Error(err)
		return
	}
}

func TestDirSize(t *testing.T) {
	s, err := os.Stat("/Users/wongkashing/Downloads/xcopy-master")
	fmt.Println(err)
	fmt.Println(s.Size())
}
