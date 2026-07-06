package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')
	texto = strings.TrimSpace(texto)

	palavras := strings.Fields(texto)

	res := palavras[0]

	for i := 1; i < len(palavras); i++ {
		ult := res[len(res)-1]
		prim := palavras[i][0]

		if vogal(ult) && vogal(prim) {
			res += palavras[i][1:]
		} else {
			res += " " + palavras[i]
		}
	}

	fmt.Println(res)
}