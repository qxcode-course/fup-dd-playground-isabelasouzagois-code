package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vet := make([]int, n)
	maior := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i])
		if vet[i] > maior {
			maior = vet[i]
		}
	}
	for nivel := maior; nivel >= 1; nivel-- {
		for i := 0; i < n; i++ {
			if vet[i] >= nivel {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}

}
