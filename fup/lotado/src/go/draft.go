package main
import "fmt"
func main() {
    var c int
    
    fmt.Scan(&c) 
    passageiro:= 0  
    for{
        var M int
        fmt.Scan(&M)
        passageiro += M

    if passageiro >=2*c{
        fmt.Println("hora de partir")
        break
    }else if passageiro ==0 {
        fmt.Println("vazio")
    }else if passageiro <c{
        fmt.Println("ainda cabe")
    }else{
         fmt.Println("lotado")
    }
    
       
    }
    
}
