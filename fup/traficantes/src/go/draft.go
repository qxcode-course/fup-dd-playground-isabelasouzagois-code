package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')
	alvo, _ := reader.ReadString('\n')
	novo, _ := reader.ReadString('\n')

	texto = strings.TrimSpace(texto)
	alvo = strings.TrimSpace(alvo)
	novo = strings.TrimSpace(novo)

	n := len(texto)
	m := len(alvo)

	var res strings.Builder

	for i := 0; i < n; {
		// se bate a palavra alvo
		if i+m <= n && texto[i:i+m] == alvo {
			res.WriteString(novo)
			i += m
		} else {
			res.WriteByte(texto[i])
			i++
		}
	}

	fmt.Println(res.String())
}