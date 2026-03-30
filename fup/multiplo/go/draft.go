package main

import "fmt"

func main() {
	var valor int
	fmt.Scan(&valor)
	if valor%7 == 0 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}

}
