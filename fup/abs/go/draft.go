package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	diferenca := num1 - num2

	if diferenca < 0 {
		diferenca = diferenca * (-1)
	}
	fmt.Printf("%d\n", diferenca)
}
