package main

import "fmt"
type Endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}
type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}
func imprimaPessoa(x Pessoa) {
	fmt.Println("Nome:", x.nome)
	fmt.Println("Idade:", x.idade)
	fmt.Println("Rua:", x.endereco.rua)
	fmt.Println("Numero:", x.endereco.numero)
	fmt.Println("Cidade:", x.endereco.cidade)
	fmt.Println("Estado:", x.endereco.estado)
}
func main() {
	x := Pessoa{nome: "Cleber",
		idade: 15,
		endereco: Endereco{
			rua:    "asa sul",
			numero: 409,
			cidade: "Brasilia",
			estado: "DF"},
	}
	imprimaPessoa(x)
}

