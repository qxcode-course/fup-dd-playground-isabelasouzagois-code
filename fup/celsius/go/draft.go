package main

import "fmt"

func main() {
	var Celsius, fahrenheit float64
	fmt.Scan(&Celsius)
	fahrenheit = 1.8*Celsius + 32
	fmt.Printf("%.6f\n", fahrenheit)
}
