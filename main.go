package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type cuncurrentMap struct {
	Map   map[string]int
	mutex chan int
}

func (c cuncurrentMap) get(key string) (int, bool) {
	v, has := c.Map[key]
	return v, has
}

func (c cuncurrentMap) set(key string, val int) {

	c.Map[key] = val
}

func main() {
	var wg sync.WaitGroup

	m := cuncurrentMap{
		Map:   make(map[string]int),
		mutex: make(chan int, 2),
	}

	m.set("one", 1)
	v, _ := m.get("one")
	fmt.Println(v)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")
			m.mutex <- 1
			fmt.Println("Lock write")
			m.set(key, i)
			<-m.mutex
		}
	}()

	wg.Add(1)
	go func() {
		// t := time.NewTicker(time.Second * 2)
		// <-t.C

		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")
			m.mutex <- 1
			fmt.Println("Lock read")
			v, _ := m.get(key)
			<-m.mutex
			fmt.Println(v)
		}
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for i := 1; i < 10; i++ {
	// 		key := strings.Join([]string{"key 2", strconv.Itoa(i)}, " ")
	// 		m.set(key, i)
	// 	}
	// }()

	wg.Wait()
}
