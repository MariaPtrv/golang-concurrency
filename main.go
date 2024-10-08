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
	ch1 := make(chan int, 10)
	// ch1 <- 1
	// ch1 <- 2
	// ch1 <- 3
	// close(ch1)

	ch2 := make(chan int, 10)
	// two <- 1
	// two <- 2
	// two <- 3
	// two <- 4
	// two <- 5
	// close(two)

	ch3 := make(chan int, 10)
	// three <- 1
	// three <- 2
	//close(three)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)
	defer close(sigCh)

	wg = sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case _, ok := <-ch1:
				if !ok {
					fmt.Println("ch1 closed")
					return
				}

			case _, ok := <-ch2:
				if !ok {
					fmt.Println("ch2 closed")
					return
				}

			case _, ok := <-ch3:
				if !ok {
					fmt.Println("ch3 closed")
					return
				}

			case <-sigCh:
				{
					fmt.Println("done")
					return
				}

			}
		}
	}()
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			select {
			default:
				// i := rand.Int()
				// ch <- i
				i1 := rand.Int()
				ch1 <- i1
				i2 := rand.Int()
				ch2 <- i2
				i3 := rand.Int()
				ch3 <- i3

			}
		}
	}()

	wg.Wait()
	close(ch1)
	close(ch2)
	close(ch3)

}
