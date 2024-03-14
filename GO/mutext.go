package main

import (
	"fmt"
	"sync"
)

func incr(x *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	*x++
}

func dcr(x *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	*x--
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	x := 8
	wg.Add(1)
	go incr(&x, &mu, &wg)
	wg.Add(1)
	go dcr(&x, &mu, &wg)

	fmt.Println(x)
	wg.Wait()
}
