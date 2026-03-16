package main

import "fmt"

func main() {
	var p1, p2, p3 float64
	var vp1, vp2, vp3 float64
	var dinheiro float64
	fmt.Scan(&p1)
	fmt.Scan(&p2)
	fmt.Scan(&p3)
	fmt.Scan(&vp1)
	fmt.Scan(&vp2)
	fmt.Scan(&vp3)
	fmt.Scan(&dinheiro)
	vp1 = vp1 * p1
	vp2 = vp2 * p2
	vp3 = vp3 * p3
	var troco float64 = dinheiro - (vp1 + vp2 + vp3)
	fmt.Printf("%.2f\n", troco)

}
