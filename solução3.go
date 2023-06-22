package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
	area   float64
}
func areaTriangulo(x Triangulo) (area float64) {
	area = (x.base * x.altura) / 2
	return area
}
func main() {
	x := Triangulo{base: 50, altura: 2}
	fmt.Println(areaTriangulo(x))
}
