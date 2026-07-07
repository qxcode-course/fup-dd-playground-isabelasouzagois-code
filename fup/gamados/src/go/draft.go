package main

import (
	"bufio"
	"fmt"
    "os"
    "strings"
)
func main() {
    reader:=bufio.NewReader(os.Stdin)
    frase, _:=reader.ReadString('\n')
    frase = strings.TrimSpace(frase)
    palavras:=strings.Fields(frase)
    for i:=0;i<len(palavras)-1; i++{
        if palavras[i]>palavras[i+1]{
            fmt.Println("nao")
            return
        }
    }
    fmt.Println("sim")
}