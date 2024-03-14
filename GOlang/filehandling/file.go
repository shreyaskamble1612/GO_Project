package main

import (
	"fmt"
	"io"
	"os"
)

func readfile(filename string) {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)

	}
	fmt.Println(string(data))
}
func main() {
	myfile, err := os.Create("test1.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(myfile.Name())

	text := "Helllo worls shreays here"
	len, err := io.WriteString(myfile, text)
	if err != nil {
		panic(err)

	}
	fmt.Println(len)

	defer myfile.Close()
	readfile(myfile.Name())

}
