package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	wg := sync.WaitGroup{}

	d1 := time.Now().Add(2 * time.Second)
	ctx1, cancel1 := context.WithDeadline(ctx, d1)

	wg.Add(1)
	go One(ctx1, &wg)

	d2 := time.Now().Add(4 * time.Second)
	ctx2, cancel2 := context.WithDeadline(ctx, d2)

	wg.Add(1)
	go Two(ctx2, &wg)

	cancel()
	fmt.Println("Cancel")
	wg.Wait()
	// time.Sleep(time.Second * 4)

	cancel2()
	fmt.Println("Cancel2")
	cancel1()
	fmt.Println("Cancel1")

}

func One(ctx context.Context, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			wg.Done()
		}
	}()
	panic("panic")
	select {
	case <-ctx.Done():
		fmt.Println("Done 1")
	}

	wg.Done()

}

func Two(ctx context.Context, wg *sync.WaitGroup) {

	select {
	case <-ctx.Done():
		fmt.Println("Done 2")
	}

	wg.Done()
}

// var wg sync.WaitGroup
// ch := make(chan int)
// sigCh := make(chan os.Signal, 1)
// defer close(sigCh)
// signal.Notify(sigCh, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM)

// wg = sync.WaitGroup{}

// wg.Add(1)
// go func() {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case v, ok1 := <-ch:
// 			if !ok1 {
// 				fmt.Println("END")
// 				return
// 			}
// 			fmt.Println(v)

// 		case <-sigCh:
// 			fmt.Println("DONE")
// 			return
// 		}
// 	}
// }()

// go func() {
// 	for {
// 		time.Sleep(500 * time.Millisecond)
// 		select {
// 		default:
// 			i := rand.Int()
// 			ch <- i
// 			// i2 := rand.Int()
// 			// ch2 <- i2
// 			// i3 := rand.Int()
// 			// ch3 <- i3
// 			// i4 := rand.Int()
// 			// ch4 <- i4
// 		}
// 	}
// }()

// // for i := 0; i < 5; i++ {
// // 	ch <- i
// // }
// // sig <- struct{}{}

// wg.Wait()
// close(ch)
