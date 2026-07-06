package main

import "fmt"

type Jogada struct {
	p1, p2 int
}

func calc_pontuacao(jogada Jogada) (bool, int) {
	if jogada.p1 < 10 || jogada.p2 < 10 {
		return false, 0
	}

	pontos := jogada.p1 - jogada.p2
	if pontos < 0 {
		pontos = -pontos
	}

	return true, pontos
}

func procurar_melhor_jogada(jogadas []Jogada) int {
	melhor := -1
	menorPontuacao := 101

	for i := 0; i < len(jogadas); i++ {
		valida, pontos := calc_pontuacao(jogadas[i])

		if valida {
			if melhor == -1 || pontos < menorPontuacao {
				melhor = i
				menorPontuacao = pontos
			}
		}
	}

	return melhor
}

func main() {
	var n int
	fmt.Scan(&n)

	jogadas := make([]Jogada, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&jogadas[i].p1, &jogadas[i].p2)
	}

	ganhador := procurar_melhor_jogada(jogadas)

	if ganhador == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(ganhador)
	}
}