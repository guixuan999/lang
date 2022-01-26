/* demonstrate the behavior of multi-goroutines iterate on the single channel
 * copyright Gui Xuan 2022
 * try it: go run goroutines_iterate_on_one_chan.go
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	natures := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			natures <- i
			time.Sleep(time.Millisecond * 200)
		}
		close(natures)
	}()

	var wg sync.WaitGroup

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(i int) {
			for item := range natures {
				fmt.Printf("coroutine[%d] got %d\n", i, item)
			}
			wg.Done()
			fmt.Printf("coroutine[%d] done\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Printf("Quiting...\n")
}
