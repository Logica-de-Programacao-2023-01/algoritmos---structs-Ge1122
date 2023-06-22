package main

import "fmt"

type Funcionario struct {
	nome    string
	idade   int
  salario float64
}

func incremento(p Funcionario, percentual float64) Funcionario {
	p.salario = p.salario + (p.salario * percentual)
	return p
}

func decremento(p Funcionario, percentual float64) Funcionario {
	p.salario = p.salario - (p.salario * percentual)
	return p
}

func calcularTempoServico(p Funcionario) int {
	return p.idade - 20
}

func main() {
	f := Funcionario{
		nome:    "Paulo",
		salario: 5000,
		idade:   25,
	}
	fmt.Println(f)
}
