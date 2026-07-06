package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mat := make([][]int, n)
	somaLinha := make([]int, n)
	somaColuna := make([]int, n)

	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&mat[i][j])
			somaLinha[i] += mat[i][j]
			somaColuna[j] += mat[i][j]
		}
	}

	maior := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			peso := somaLinha[i] + somaColuna[j] - 2*mat[i][j]
			if peso > maior {
				maior = peso
			}
		}
	}

	fmt.Println(maior)
}