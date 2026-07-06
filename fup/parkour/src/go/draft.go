package main
import "fmt"
func main() {
   

    var elementos int

    fmt.Scan(&elementos)

    vetorElementos := make([]int, elementos)

    for i := 0; i < elementos; i++ {
        fmt.Scan(&vetorElementos[i])
    }

    movimentos := 0

    for i := 0; i < elementos-1; i++ {
        if vetorElementos[i] < vetorElementos[i+1] && vetorElementos[i+1]-vetorElementos[i] > 1 {
            movimentos++
        } else if vetorElementos[i] > vetorElementos[i+1] && vetorElementos[i]-vetorElementos[i+1] > 1 {
            movimentos++
        }
    }

    fmt.Println(movimentos)

}