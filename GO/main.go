// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// 	// "sync"
// 	// "time"
// )

// func printnum(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(100 * time.Millisecond)
// 	}

// }

// func printchar(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 'a'; i <= 'e'; i++ {
// 		fmt.Println(string(i))
// 		time.Sleep(100 * time.Millisecond)
// 	}

// }
// func main() {
// 	// fmt.Print("Hello")v
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go printnum(&wg)
// 	wg.Add(1)
// 	go printchar(&wg)

// 	wg.Wait()
// 	// time.Sleep(1500 * time.Millisecond)

// }
