package main

import (
	"fmt"
)

type Player struct {
	firstname, lastname string
	country             string
	age                 int
}

func main() {
	// var player1 = Player("shreyas", "kamble", "india", 30)
	var player1 = Player(firstname: "shreyas",lastname:"kambnle",country:"india",age:20)
	player1.display()
	// fmt.Println(player1)

}
