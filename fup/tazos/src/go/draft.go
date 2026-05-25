package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	maiorFreq := 1

	freq := 1

	resultado := []int{}

	for i := 1; i <= n; i++ {

		if i < n && vetor[i] == vetor[i-1] {
			freq++
		} else {

			// encontrou frequência maior
			if freq > maiorFreq {
				maiorFreq = freq
				resultado = []int{vetor[i-1]}
			} else if freq == maiorFreq {
				resultado = append(resultado, vetor[i-1])
			}

			freq = 1
		}
	}

	fmt.Print("[ ")

	for i := 0; i < len(resultado); i++ {
		fmt.Print(resultado[i], " ")
	}

	fmt.Println("]")
}