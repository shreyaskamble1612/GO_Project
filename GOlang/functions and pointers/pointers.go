package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var myptr *int
	myptr = &a
	*myptr = 15
	// fmt.Println(myptr)
	fmt.Println(*myptr)

}
