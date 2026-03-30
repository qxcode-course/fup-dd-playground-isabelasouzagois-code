package main

import "fmt"

func main() {
	var d1, d2, d3 int
	fmt.Scan(&d1, &d2, &d3)
	soma := d1 + d2 + d3
	if soma%2 == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(1)

	}
}
