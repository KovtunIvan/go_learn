package channels

import (
	"fmt"
	"sync"
)

func Example() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 1
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()
	wg.Wait()
	// 1
}
