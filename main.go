package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)

	wg = sync.WaitGroup{}

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

			case <-sigCh:
				fmt.Println("DONE")
				return
			}
		}
	}()

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			select {
			default:
				i := rand.Int()
				ch <- i
				// i2 := rand.Int()
				// ch2 <- i2
				// i3 := rand.Int()
				// ch3 <- i3
				// i4 := rand.Int()
				// ch4 <- i4
			}
		}
	}()

	// for i := 0; i < 5; i++ {
	// 	ch <- i
	// }
	// sig <- struct{}{}

	wg.Wait()
	close(ch)

}
