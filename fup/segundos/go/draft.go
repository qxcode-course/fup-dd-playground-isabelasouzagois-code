package main

import "fmt"

func main() {
	var hora, minutos, segundos int
	fmt.Scan(&segundos)
	hora = segundos / 3600

	minutos = (segundos % 3600) / 60
	segundos = (segundos % 3600) % 60
	fmt.Printf("%d:%d:%d\n", hora, minutos, segundos)
}
