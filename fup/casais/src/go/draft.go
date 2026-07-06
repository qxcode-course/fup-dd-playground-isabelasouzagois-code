package main
import "fmt"
func casal(elemento int)int{
    vetor:= make([]int,elemento)
    for i:=0;i<elemento;i++{
        fmt.Scan(&vetor[i])
    }
    count:= 0
    for i:=0;i< elemento - 1 ;i++{
        
            if (vetor[i] * -1) == vetor[i + 1]{
                count++
            }
        
        
}  
 return count
}
func main() {
    var elemento int
    fmt.Scan(&elemento)
    fmt.Print(casal(elemento))

}
