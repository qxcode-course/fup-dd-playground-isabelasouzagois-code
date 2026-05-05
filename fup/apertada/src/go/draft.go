package main
import "fmt"
func main() {
    var v [5]int
    
    for i:=0;i<5;i++{
        fmt.Scan(&v[i])
        
        }
        menor:=v[0]
         for i:=1;i<5;i++{
       if v[i]<menor{
        menor=v[i]
       }
    }
    fmt.Println(menor)
}
