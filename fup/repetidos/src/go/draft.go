package main
import "fmt"
func main() {
    var numeros, procurado int
    count := 0
    fmt.Scan(&procurado, &numeros)

    var vetor = make([]int, numeros)

    for i := 0; i < numeros; i++ {
        fmt.Scan(&vetor[i])
    }

    for i := 0; i < numeros; i++ {
       if(vetor[i] == procurado){
        count += 1
       } 
    }
    fmt.Print(count)
    fmt.Printf("\n")
}
