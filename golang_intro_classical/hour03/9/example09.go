package main

import "fmt"

var s = "Hello World"

func main() {
	fmt.Printf("Printing 's' variable from outer block %v\n", s)
	b := true
	if b {
		fmt.Printf("Printing 'b' variable from outer block %v\n", b)
		i := 42
		if b != false {
			fmt.Printf("Printing 'i' variable form outer block %v\n", i)
		}
	}
}
