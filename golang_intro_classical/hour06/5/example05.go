package main

import "fmt"

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgogne"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	//copy(smellyCheeses, cheeses[1:])
	fmt.Println(smellyCheeses)
}
