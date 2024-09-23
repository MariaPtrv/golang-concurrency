package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	ch2 := make(chan int, 3)
	ch2 <- 4
	ch2 <- 5
	ch2 <- 6

	go print(ch1)
	go print(ch2)
	wg.Wait()
	time.Sleep(time.Second * 2)
	go print(merge(ch1, ch2))
}

func merge(chs ...chan int) chan int {
	chRes := make(chan int, 6)
	for _, ch := range chs {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				chRes <- v
			}
		}(ch)
	}

	return chRes
}

func fillCh(ch chan int) chan int {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

func print(ch chan int) {
	fmt.Println("chan ---")
	for v := range ch {
		fmt.Printf("%d ", v)
	}
	fmt.Println(" --- chan")
}
