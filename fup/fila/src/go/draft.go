package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	par := make([]int, n)
	impar := make([]int, n)

	qtdPar := 0
	qtdImpar := 0

	for i := 0; i < n; i++ {
		if vetor[i]%2 == 0 {
			par[qtdPar] = vetor[i]
			qtdPar++
		} else {
			impar[qtdImpar] = vetor[i]
			qtdImpar++
		}
	}

	fmt.Print("[ ")
	for i := 0; i < qtdImpar; i++ {
		fmt.Printf("%d ", impar[i])
	}
	fmt.Println("]")

	fmt.Print("[ ")
	for i := 0; i < qtdPar; i++ {
		fmt.Printf("%d ", par[i])
	}
	fmt.Println("]")
}