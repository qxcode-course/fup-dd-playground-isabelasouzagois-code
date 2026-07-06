package main

import "fmt"

type Restaurante struct {
	nome  string
	ponto int
}

func melhorRestaurante(lista []Restaurante) string {
	melhor := lista[0]

	for i := 1; i < len(lista); i++ {
		if lista[i].ponto > melhor.ponto {
			melhor = lista[i]
		} else if lista[i].ponto == melhor.ponto {
			if lista[i].nome < melhor.nome {
				melhor = lista[i]
			}
		}
	}

	return melhor.nome
}

func main() {
	var n int
	fmt.Scan(&n)

	restaurantes := make([]Restaurante, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&restaurantes[i].nome, &restaurantes[i].ponto)
	}

	fmt.Println(melhorRestaurante(restaurantes))
}
