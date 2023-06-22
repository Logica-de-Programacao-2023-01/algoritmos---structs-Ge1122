package main

import "fmt"

type Musica struct {
	titulo  string
	duracao int
  artista string
}

func encontrarMusicas(playlist []Playlist, nomeMusica string) []Playlist {
	var resultado []Playlist
	for _, playlist := range playlist {
		for _, musica := range playlist.musicas {
			if musica.titulo == nomeMusica {
				resultado = append(resultado, playlist)	}
		}
	}
	return resultado
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func main() {
	p := Playlist{
		nome: "Antonio",
		musicas: []Musica{
			{
				titulo:  "mercy",
				artista: "Shawn Mendes",
				duracao: 3.28,
			},
    {
				titulo:  "tell me you love me",
				artista: "Demi Lovato",
				duracao: 3.56,
			},
		},
	}
	fmt.Println(p)
}
