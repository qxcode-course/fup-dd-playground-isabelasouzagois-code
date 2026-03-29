package main

import "fmt"

func main() {
	var num1 int
	fmt.Scan(&num1)

	if num1 >= 0 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}

}
