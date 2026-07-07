package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	mat := make([][]string, n)

	linhaLeao, colunaLeao := -1, -1

	for i := 0; i < n; i++ {
		mat[i] = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&mat[i][j])
			if mat[i][j] == "L" {
				linhaLeao = i
				colunaLeao = j
			}
		}
	}

	gladiadores := 0
	condenados := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			
			if i == linhaLeao || j == colunaLeao {
				continue
			}

			if mat[i][j] == "G" {
				gladiadores += 2
			} else if mat[i][j] == "C" {
				if i+j == n-1 {
					condenados += 2
				} else {
					condenados++
				}
            }

		}
	}

	if gladiadores > condenados {
		fmt.Println("Gladiadores")
	} else if condenados > gladiadores {
		fmt.Println("Condenados a morte")
	} else {
		fmt.Println("Ninguem")
	}
}