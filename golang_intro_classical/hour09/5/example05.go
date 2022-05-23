package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 1
	var s string = " egg"
	intToString := strconv.Itoa(i)
	var breakfast string = intToString + s
	fmt.Println(breakfast)
}
