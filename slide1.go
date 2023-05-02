package main

import (
	"fmt"
)

type medidas struct {
	altura float64
	base   float64
}

func main() {
	a := medidas{altura: 10, base: 10}
	fmt.Println("Qual a altura? ")
	fmt.Scan(&a.altura)
	fmt.Println("Qual a base? ")
	fmt.Scan(&a.base)
	fmt.Println(area(a))

}
func area(a medidas) float64 {
	return a.altura * a.base
}
