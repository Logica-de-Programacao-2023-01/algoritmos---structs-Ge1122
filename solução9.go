package main

import "fmt"
type Estudante struct {
	nome  string
	idade int
	notas []float64
}

func (x Estudante) adicionarNota(nota float64) Estudante {
	x.notas = append(x.notas, nota)
	return x
}

func (x Estudante) removerNota(indice int) Estudante {
	if indice >= 0 && indice < len(x.notas) {
		x.notas = append(x.notas[:indice], x.notas[indice+1:]...)
	}
	return x
}

func (x Estudante) calculaMedia() float64 {
	total := 0.0
	for _, nota := range x.notas {
		total += nota
	}
	if len(x.notas) > 0 {
		return total / float64(len(x.notas))
	}
	return 0.0
}

func (x Estudante) imprimirInformacao() {
	fmt.Println("Nome:", x.nome)
	fmt.Println("Media:", x.calculaMedia())
  fmt.Println("Idade:", x.idade)
	
}

func main() {
	x := Estudante{
		nome:  "Georges",
		idade: 18,
		notas: []float64{9.5, 8.9, 10},
	}
	x = x.adicionarNota(8.9)
	x.imprimirInformacao()

	x = x.removerNota(1)
	x.imprimirInformacao()
}
