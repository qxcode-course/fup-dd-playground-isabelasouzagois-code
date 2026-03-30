package main

import "fmt"

func main() {
	var idade int
	fmt.Scan(&idade)

	if idade < 12 {
		fmt.Println("crianca")
	} else if idade < 18 {
		fmt.Println("jovem")
	} else if idade < 65 {
		fmt.Println("adulto")
	} else if idade < 1000 {
		fmt.Println("idoso")
	} else {
		fmt.Println("mumia")
	}
}
