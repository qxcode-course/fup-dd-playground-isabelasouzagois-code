package main

import "fmt"

func main() {
	cartela := [4][4]int{
		{1,  9,  27, 23},
		{34, 20, 37, 47},
		{30, 87, 55, 69},
		{13, 60, 99, 66},
	}

	cont := 0

	for i := 0; i < 6; i++ {
		var num int
		fmt.Scan(&num)

		for l := 0; l < 4; l++ {
			for c := 0; c < 4; c++ {
				if num == cartela[l][c] {
					cont++
				}
			}
		}
	}

	fmt.Println(cont)
}