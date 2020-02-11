package main

import (
	"fmt"
	"sync"
)

// 锁

var x int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("now x variable is:", x)
}
