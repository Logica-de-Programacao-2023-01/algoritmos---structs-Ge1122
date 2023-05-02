package main

import "fmt"

type aluno struct {
	nome  string
	idade int
	notas []float64
}

func media(a aluno) float64 {
	var soma float64
	for _, nota := range a.notas {
		soma += nota
	}
	return soma / float64(len(a.notas))
}
func main() {
	a := aluno{
		nome:  "pedro",
		idade: 13,
		notas: []float64{1, 1, 1, 1, 1},
	}
	fmt.Println(media(a))
}
