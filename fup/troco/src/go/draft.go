package main
import "fmt"
func main() {
    var troco float64
    fmt.Scan(&troco)
    vetor:=[]float64{100,50, 20, 10, 5, 2,1.00, 0.50, 0.25,0.10, 0.05}
    cedulas:= make([]int,vetor)
    for i:=0;i<len(vetor);i++{
       if troco>=vetor[i]{
        cedulas[i]= int(troco/(vetor[i]))
       }
    }
    fmt.Println("Hello, World!")
}