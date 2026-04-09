

package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	dx := x2 - x1
	dy := y2 - y1
	soma := dx*dx + dy*dy
	distancia := math.Sqrt(soma)
	fmt.Printf("%.2f\n", distancia)
}
