package main

import "fmt"

type Fator struct {
	num int
	qtd int
}

func calc_fatores(num int) []Fator {
	var fatores []Fator

	divisor := 2

	for num > 1 {
		cont := 0

		for num%divisor == 0 {
			cont++
			num = num / divisor
		}

		if cont > 0 {
			f := Fator{divisor, cont}
			fatores = append(fatores, f)
		}

		divisor++
	}

	return fatores
}

func main() {
	var n int
	fmt.Scan(&n)

	fatores := calc_fatores(n)

	for i := 0; i < len(fatores); i++ {
		fmt.Println(fatores[i].num, fatores[i].qtd)
	}
} 