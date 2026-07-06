package main

import "fmt"

func main() {
	var mat [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&mat[i][j])
		}
	}

	soma := 0

	// soma da primeira linha como referência
	for j := 0; j < 3; j++ {
		soma += mat[0][j]
	}

	// verifica linhas
	for i := 1; i < 3; i++ {
		atual := 0
		for j := 0; j < 3; j++ {
			atual += mat[i][j]
		}
		if atual != soma {
			fmt.Println("nao")
			return
		}
	}

	// verifica colunas
	for j := 0; j < 3; j++ {
		atual := 0
		for i := 0; i < 3; i++ {
			atual += mat[i][j]
		}
		if atual != soma {
			fmt.Println("nao")
			return
		}
	}

	// diagonal principal
	atual := mat[0][0] + mat[1][1] + mat[2][2]
	if atual != soma {
		fmt.Println("nao")
		return
	}

	// diagonal secundária
	atual = mat[0][2] + mat[1][1] + mat[2][0]
	if atual != soma {
		fmt.Println("nao")
		return
	}

	fmt.Println("sim")
}