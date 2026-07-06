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
	padrao, _ := reader.ReadString('\n')

	texto = strings.TrimSpace(texto)
	padrao = strings.TrimSpace(padrao)

	n := len(texto)
	m := len(padrao)

	cont := 0

	for i := 0; i <= n-m; i++ {
		if texto[i:i+m] == padrao {
			cont++
		}
	}

	fmt.Println(cont)
}