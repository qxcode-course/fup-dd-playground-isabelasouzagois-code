package main

import "fmt"

func main() {
	var chico, cebolinha int
	var n int

	fmt.Scan(&chico)
	fmt.Scan(&cebolinha)
	fmt.Scan(&n)

	totalPatas := 0

	for i := 0; i < n; i++ {
		var animal string
		fmt.Scan(&animal)

		if animal == "v" || animal == "c" {
			totalPatas += 4
		} else if animal == "g" {
			totalPatas += 2
		}
	}

	fmt.Println(totalPatas)

	difChico := chico - totalPatas
	if difChico < 0 {
		difChico = -difChico
	}

	difCebolinha := cebolinha - totalPatas
	if difCebolinha < 0 {
		difCebolinha = -difCebolinha
	}

	if difChico < difCebolinha {
		fmt.Println("Chico Bento")
	} else if difCebolinha < difChico {
		fmt.Println("Cebolinha")
	} else {
		fmt.Println("empate")
	}
}