package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func main() {
	sigCh := make(chan os.Signal, 1)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)

	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	ch2 := make(chan int, 3)
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6

	// go print(ch1)
	// go print(ch2)
	//time.Sleep(time.Second * 2)

	sig := make(chan struct{})
	ch := merge(sig, ch1, ch2)

	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println(v)
		case <-time.After(time.Second * 10):
			sig <- struct{}{}
			fmt.Println("---")
		case <-sigCh:
			fmt.Println("cancel")
			sig <- struct{}{}
			return
		}

	}

}

func merge(sig chan struct{}, chs ...chan int) chan int {
	chRes := make(chan int)
	for _, ch := range chs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case temp, ok := <-ch:
					if !ok {
						return
					}
					chRes <- temp

				case _ = <-sig:
					return
				}
			}
		}()

		// go func(ch chan int) {
		// 	for v := range ch {
		// 		chRes <- v
		// 	}
		// 	wg.Done()
		// 	fmt.Println("after done")
		// 	return
		// }(ch)
	}

	go func() {
		fmt.Println("before")
		wg.Wait()
		fmt.Println("after wait")
		close(chRes)

	}()

	return chRes
}

func fillCh(ch chan int) chan int {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

// func print(ch chan int) {
// 	//fmt.Println("chan ---")
// 	for v := range ch {
// 		//fmt.Printf("%d ", v)
// 	}
// 	//fmt.Println(" --- chan")
// }
