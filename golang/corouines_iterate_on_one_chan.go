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
			time.Sleep(time.Millisecond * 100)
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
