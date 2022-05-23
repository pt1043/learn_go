package main

import "fmt"

func main() {

	i := 0
	for {
		fmt.Println("你好，Golang", i)
		i++
		if i >= 10 {
			break
		}
		if i%2 == 0 {
			continue
		}
		fmt.Println("jump", i)
	}
}
