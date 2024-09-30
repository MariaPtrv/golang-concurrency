package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	sig := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok1 := <-ch:
				if !ok1 {
					fmt.Println("END")
					return
				}
				fmt.Println(v)

			case <-sig:
				fmt.Println("DONE")
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
	}
	sig <- struct{}{}

	close(ch)

	wg.Wait()

}
