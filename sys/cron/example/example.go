package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/sys/cron"
	"sync"
	"time"
)

var wg sync.WaitGroup

func tryFreq(fm *cron.RateLimiter, id int) {
	defer wg.Add(-1)
	for i := 0; i < 5; i++ {
		fm.MarkAndWaitBlock()
		fmt.Println(fmt.Sprintf("Routine %d got a frequency mutex", id))
	}
}

func main() {
	fm := cron.NewRateLimiter(time.Millisecond * 1000)

	count := 10
	wg.Add(count)
	for i := 0; i < count; i++ {
		go tryFreq(fm, i)
	}
	wg.Wait()
}
