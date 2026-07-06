package main

import "fmt"

func overlap(a, b string) int {
	lim := len(a)
	if len(b) < lim {
		lim = len(b)
	}

	for i := lim; i >= 1; i-- {
		if a[len(a)-i:] == b[:i] {
			return i
		}
	}
	return 0
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	for {
		k := overlap(a, b)
		if k == 0 {
			break
		}
		a = a[:len(a)-k]
		b = b[k:]
	}

	fmt.Println(a + b)
}