package main

import "fmt"

func main() {
	// names := make([]int,)
	names := []int{1, 25, 98, 6, 5}
	fmt.Println(names)
	names = append(names, 1, 25, 98)
	fmt.Println(names)
}
