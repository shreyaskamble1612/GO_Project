// package main

// import "fmt"

// func sum(a, b int, ch chan int) {
// 	c := a + b
// 	ch <- c
// }

// func main() {
// 	ch := make(chan int)

// 	go sum(4, 5, ch)
// 	output := <-ch

// 	fmt.Println(output)

// }
