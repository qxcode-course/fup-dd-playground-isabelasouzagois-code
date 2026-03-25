package main

import "fmt"

func main() {
	var primeiro, segundo int
	var op string

	fmt.Scan(&primeiro, &segundo, &op)
	if op == "+" {
		fmt.Println(primeiro + segundo)

	} else if op == "-" {
		fmt.Println(primeiro - segundo)

	} else if op == "*" {
		fmt.Println(primeiro * segundo)

	} else {
		fmt.Println(primeiro / segundo)
	}

}
