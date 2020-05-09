package main

import (
	"fmt"
	"github.com/shawnwyckoff/gopkg/apputil/ginstance"
	"log"
	"time"
)

func main() {
	lock := ginstance.New("your-app-name")
	defer lock.UnLock()

	ok, err := lock.IsSingle()
	if err != nil {
		fmt.Println(err)
		return
	}
	if ok {
		fmt.Println("Is single process")
	} else {
		fmt.Println("Another process running")
	}

	time.Sleep(60 * time.Second)
	log.Println("finished")
}
