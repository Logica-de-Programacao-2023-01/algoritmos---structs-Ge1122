package main

import "fmt"
type musicas struct {
	titulo  string
	artista string
	duracao float64
}
type Playlists struct {
	nome   string
	musica []musicas
}

func imprimaPlaylist(m Playlists) {
	fmt.Println("Nome da Playlist:", m.nome)
	var soma float64 = 0
	for _, a := range m.musica {
		soma += a.duracao
		fmt.Println("Duracao total:", soma)
		fmt.Println("Titulo:", m.titulo)
		fmt.Println("Artista:", m.artista)
		fmt.Println("Duracao:", m.duracao)
		fmt.Println("")
	}
}

func main() {
	m := Playlists{
		nome: "pop",
		musica: []musicas{
			{
				titulo:  "mercy",
				artista: "Shawn Mendes",
				duracao: 3.28,
			},
			{
				titulo:  "Pumped Up Kicks",
				artista: "Madism, MKJ, Felix Samuel",
				duracao: 2.25,
			},
		},
	}
	imprimaPlaylist(m)

}

