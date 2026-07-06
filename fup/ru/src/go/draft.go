package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	frase, _ := reader.ReadString('\n')

	vogais := ""
	consoantes := ""

	for _, c := range frase {
		if c == ' ' || c == '\n' {
			continue
		}

		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vogais += string(c)
		} else {
			consoantes += string(c)
		}
	}

	fmt.Println(vogais)
	fmt.Println(consoantes)
}