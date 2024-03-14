package main

import (
	"fmt"
)

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
func main() {
	a := 10
	b := 20
	fmt.Println(a, b)
	swap2(&a, &b)

	fmt.Println(a, b)
}
