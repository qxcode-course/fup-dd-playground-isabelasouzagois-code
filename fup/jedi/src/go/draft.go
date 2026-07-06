package main
import "fmt"
func main() {
    var nElementos int

    fmt.Scan(&nElementos)

    vetor := make([]int, nElementos)

    for i := 0; i < nElementos; i++ {
        fmt.Scan(&vetor[i])
    }

    forcaJ := 0
    forcaS := 0

    for i := 0; i < nElementos; i++ {
        if i < (nElementos / 2) {
            forcaJ += vetor[i]
        } else if i >= (nElementos / 2) {
            forcaS += vetor[i]
        }
    }

    if forcaJ > forcaS {
        fmt.Println("Jedi")
    } else if forcaS > forcaJ {
        fmt.Println("Sith")
    } else if forcaS == forcaJ {
        fmt.Println("Empate")
    }
}