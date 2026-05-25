package main

import (
	"bufio"
	"fmt"
	"os"
	
)
func main() {
    scanner :=bufio.NewScanner(os.Stdin)
    scanner.Scan()
    texto:=scanner.Text()
    runes:=[]rune(texto)
    for i:=len(runes)-1;i>=0;i--{
      fmt.Print(string(runes[i]))
      
    }
    fmt.Printf("\n")
}
