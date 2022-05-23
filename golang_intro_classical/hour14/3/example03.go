package main

import "fmt"

const Foo string = "constant string"

func main() {
	fmt.Println(Foo)
	a_string := "hello"
	fmt.Println(a_string)
}

//go get -u github.com/golang/lint/golint
//go lint example02.go
