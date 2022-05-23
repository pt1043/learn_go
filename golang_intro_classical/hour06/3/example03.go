package main

import "fmt"

func main() {
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Epoisses de Bourgogne"
	cheeses = append(cheeses, "camembert", "Reblochon", "Picodon")
	fmt.Println(cheeses)
}
