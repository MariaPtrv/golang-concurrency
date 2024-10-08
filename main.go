package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type cuncurrentMap struct {
	Map map[string]int
	mu  sync.Mutex
}

func (c cuncurrentMap) get(key string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, has := c.Map[key]
	return v, has
}

func (c cuncurrentMap) set(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Map[key] = val
}

func main() {
	var wg sync.WaitGroup

	m := cuncurrentMap{
		Map: make(map[string]int),
	}

	m.set("one", 1)
	v, _ := m.get("one")
	fmt.Println(v)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")
			m.set(key, i)
		}
	}()

	wg.Add(1)
	go func() {
		t := time.NewTicker(time.Second * 2)
		<-t.C

		defer wg.Done()
		for i := 1; i < 10; i++ {
			key := strings.Join([]string{"key", strconv.Itoa(i)}, " ")
			v, _ := m.get(key)
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
