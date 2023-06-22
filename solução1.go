package main
import (
	"fmt"
	"math"
)
type Circulo struct {
	raio float64
	area float64
}
func areaCirculo(c Circulo) (area float64) {
	area = math.Pi * c.raio * c.raio
	return area
}
func main() {
	c := Circulo{raio: 5}
	fmt.Println(areaCirculo(c))
}

