package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scan(&num1, &num2)

	if num1 == 5 || num1 == 3 || num2 == 5 || num2 == 3 {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}

}
