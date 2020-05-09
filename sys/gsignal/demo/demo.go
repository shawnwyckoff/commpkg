package main

import (
	"fmt"
	"github.com/shawnwyckoff/gopkg/xsystils/xsys/xexit"
	"github.com/shawnwyckoff/gopkg/xsystils/xsys/xfs"
	"os"
	"time"
)

func exitcb(sig os.Signal, closemsg string) {
	s := fmt.Sprintf("exit signal %s, closemsg %s", sig.String(), closemsg)
	fmt.Println(s)
	xfs.AppendStringToFile(s, "demo.log")
}

func main() {
	xexit.RegisterExitCallback(exitcb)
	/*go func() {
		type obj struct {
			name string
		}
		o := new(obj)
		o = nil
		o.name = ""
		time.Sleep(time.Second * 5)
		os.Exit(-1)
	}()*/

	for {
		time.Sleep(time.Second)
	}
}
