package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a)
	fmt.Scan(&b)
	var media float32 = (a + b) / 2
	fmt.Printf("%.1f\n",media)

}
