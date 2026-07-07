package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	machos := make(map[int]int)
	femeas := make(map[int]int)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x > 0 {
			machos[x]++
		} else {
			femeas[-x]++
		}
	}

	casais := 0

	for especie, qtdMachos := range machos {
		qtdFemeas := femeas[especie]
		if qtdMachos < qtdFemeas {
			casais += qtdMachos
		} else {
			casais += qtdFemeas
		}
	}

	fmt.Println(casais)
}