package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	w1, _ := reader.ReadString('\n')
	w2, _ := reader.ReadString('\n')

	// remove '\n' se existir
	if len(text) > 0 && text[len(text)-1] == '\n' {
		text = text[:len(text)-1]
	}
	if len(w1) > 0 && w1[len(w1)-1] == '\n' {
		w1 = w1[:len(w1)-1]
	}
	if len(w2) > 0 && w2[len(w2)-1] == '\n' {
		w2 = w2[:len(w2)-1]
	}

	// mapa de substituição (bidirecional)
	m := make(map[byte]byte)

	for i := 0; i < len(w1) && i < len(w2); i++ {
		m[w1[i]] = w2[i]
		m[w2[i]] = w1[i]
	}

	res := []byte(text)

	for i := 0; i < len(res); i++ {
		if v, ok := m[res[i]]; ok {
			res[i] = v
		}
	}

	fmt.Println(string(res))
}