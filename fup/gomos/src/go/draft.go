package main

import "fmt"

type Gomo struct {
	x int
	y int
}

func main() {
	var q int
	var d string

	fmt.Scan(&q, &d)

	cobra := make([]Gomo, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&cobra[i].x, &cobra[i].y)
	}

	
	xAnt := cobra[0].x
	yAnt := cobra[0].y

	
	if d == "L" {
		cobra[0].x--
	} else if d == "R" {
		cobra[0].x++
	} else if d == "U" {
		cobra[0].y--
	} else if d == "D" {
		cobra[0].y++
	}

	for i := 1; i < q; i++ {
		auxX := cobra[i].x
		auxY := cobra[i].y

		cobra[i].x = xAnt
		cobra[i].y = yAnt

		xAnt = auxX
		yAnt = auxY
	}

	
	for i := 0; i < q; i++ {
		fmt.Println(cobra[i].x, cobra[i].y)
	}
}