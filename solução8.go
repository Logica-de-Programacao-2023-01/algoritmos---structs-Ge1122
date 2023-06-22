package main
type Viagem struct {
	origem  string
	destino string
	data    int
	preco   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	var viagemMax Viagem
	for _, viagem := range viagens {
		if viagem.preco > viagemMax.preco {
			viagemMax = viagem
		}
	}
	return viagemMax
}

func main() {
	BSBpRJ := viagem{
		Origem:  "São Paulo",
		Destino: "Rio de Janeiro",
		Data:    10 / 06,
		Preco:   1900,
	}
	BSBpTUR := viagem{
		Origem:  "Brasília",
		Destino: "Istanbul - Turquia",
		Data:    20 / 12,
		Preco:   7100,
	}
	BSBpDF := viagem{
		Origem:  "Salvador",
		Destino: "Brasília",
		Data:    11 / 07,
		Preco:   1500,
	}

	opcoesViagem := []viagem{BSBpRJ, BSBpDF, BSBpTUR}

	fmt.Println("A viagem mais cara é:", ViagemMaisCara(opcoesViagem))
}
