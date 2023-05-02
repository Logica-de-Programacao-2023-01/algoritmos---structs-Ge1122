package main

//Crie uma struct chamada Livro com os campos "título",
//"autor" e "ano de publicação".
//Escreva uma função que receba um Livro como parâmetro e imprima suas informações.

import "fmt"

type livro struct {
	titulo  string
	autor   string
	ano     int
	titulo2 string
	autor2  string
	ano2    int
}

func main() {
	var x string
	l := livro{titulo: "Harry.potter", autor: "J.K.Rowlling", ano: 1997, titulo2: "SA", autor2: "paulin", ano2: 1923}
	fmt.Println("So temos os livros Harry,potter e SA")
	fmt.Println("Qual o livro que deseja escolher")
	fmt.Scan(&x)
	if x == l.titulo {
		fmt.Println("O titulo do livro é: ", l.titulo)
		fmt.Println("O(a) autor(a) é: ", l.autor)
		fmt.Println("A data de pub do livro é: ", l.ano)
	} else if x == l.titulo2 {
		fmt.Println("O titulo do livro é: ", l.titulo2)
		fmt.Println("O(a) autor(a) é: ", l.autor2)
		fmt.Println("A data de pub do livro é: ", l.ano2)
	}

}
