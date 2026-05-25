package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// vetor para sabores
	sabores := make([]string, n)

	// vetor para turnos
	turnos := make([]string, n) 

	// leitura dos dados
	for i := 0; i < n; i++ {
		fmt.Scan(&sabores[i], &turnos[i])
	}

	// contadores
	chocolate := 0
	limao := 0

	manha := 0
	tarde := 0

	// percorrendo os vetores
	for i := 0; i < n; i++ {

		if sabores[i] == "c" {
			chocolate++
		} else if sabores[i] == "l" {
			limao++
		}

		if turnos[i] == "m" {
			manha++
		} else if turnos[i] == "t" {
			tarde++
		}
	}

	// sabor mais vendido
	if chocolate > limao {
		fmt.Println("c")
	} else if limao > chocolate {
		fmt.Println("l")
	} else {
		fmt.Println("empate")
	}

	// turno mais vago (menos vendas)
	if manha < tarde {
		fmt.Println("m")
	} else if tarde < manha {
		fmt.Println("t")
	} else {
		fmt.Println("empate")
	}
}