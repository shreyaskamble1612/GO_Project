package main

import "fmt"

func main() {
	// map_name := make(map[int]int)

	lang := make(map[int]string)

	lang[0] = "C++"
	lang[1] = "GO"
	lang[2] = "JavaScript"
	fmt.Println(lang)
	delete(lang, 2)
	fmt.Println(lang)
}
