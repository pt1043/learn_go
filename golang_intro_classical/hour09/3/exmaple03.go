package main

import "fmt"

func main() {
	s := `After a backslash, certain single character escapes represent special values
n is aline feed or new line
	t is a tab`
	fmt.Println(s)
}
