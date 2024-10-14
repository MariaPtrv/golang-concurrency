package main

import (
  "fmt"
  "sync"
  "time"
)

func exampleRangeChannel() {
  ch := make(chan int)

  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    defer wg.Done()
    for k := range ch {
      fmt.Println(k)
    }
    fmt.Println("exit")
  }()

  for i := 1; i <= 10; i++ {
    ch <- i
  }
  fmt.Println("before close")
  close(ch)
  fmt.Println("after close")
  wg.Wait()
  fmt.Println("end")
}

func exampleSelectChannelClosed() {
  ch := make(chan int)

  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    defer wg.Done()
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
  Loop:
    for {
      select {
      // этот select будет блокировать цикл for пока канал не закроется
      case val, ok := <-ch:
        //time.Sleep(500 * time.Millisecond)
        if !ok {
          fmt.Println("done")
          break Loop
        }
        fmt.Println(val)
        // этот тикер остановит выполнение селекта, когда выйдет время и ветвь селекта case val, ok := <-ch будет разблокирована
      case <-ticker.C:
        fmt.Println("time has ended")
        break Loop
        // тут можно послушать контекст, если нужно выключить по сигналу горутину. В таком случае нужно использовать селект вместо цикла для заполнения канала 
      }
    }
    fmt.Println("exit")
  }()

  for i := 1; i <= 1; i++ {
    ch <- i
    //time.Sleep(1 * time.Second)
  }
  fmt.Println("before close")
  // срабатывание тикера в горутине
  time.Sleep(1500 * time.Millisecond)
  close(ch)
  fmt.Println("after close")
  wg.Wait()
  fmt.Println("end")
}

func main() {
  exampleRangeChannel()
  fmt.Println("-------------")
  exampleSelectChannelClosed()
}