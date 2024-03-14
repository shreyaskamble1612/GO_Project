package main

import "fmt"

// func add(a int, b int) int {
// 	var res = a + b
// 	return res
// }

// func isEven(a int) string {
// 	if a%2 == 0 {
// 		return "even"
// 	} else {
// 		return "odd"
// 	}
// }

// func arraysum(arr []int) int {
// 	var sum int
// 	for _, val := range arr {
// 		sum += val
// 	}
// 	return sum
// }

func arithmetic(a, b int) (int, int, int, float32) {
	var sum = a + b
	var sub = a - b
	var mul = a * b
	var div = float32(a) / float32(b)
	return sum, sub, mul, div
}
func main() {
	// fmt.Printf("%d is %s", 25, isEven(25))

	// arr := []int{12, 58, 65, 14, 25}
	// fmt.Println("Sum of all the elenmt of array is %d", arraysum(arr))
	add, sub, div, mul := arithmetic(5, 6)

	fmt.Println(add, sub, div, mul)
}
