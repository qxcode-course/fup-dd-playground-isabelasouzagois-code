package main
import "fmt"
func main() {
    var texto string
var inicio, qtd int

fmt.Scan(&texto, &inicio, &qtd)

if inicio < 0 || inicio >= len(texto) || qtd <= 0 {
    fmt.Println("")
    return
}

fim := inicio + qtd
if fim > len(texto) {
    fim = len(texto)
}

fmt.Println(texto[inicio:fim])


}