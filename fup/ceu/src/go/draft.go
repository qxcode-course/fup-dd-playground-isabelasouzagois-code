package main
import "fmt"
func main() {
    var pedra int
    fmt.Scan(&pedra)
    fmt.Print("[ ")
    for i := 0; i <= 10; i++{
        if i == pedra{
            continue
        }

        if i == 10{
            fmt.Print("ceu ")
        
        }else{
            fmt.Printf("%d ",i)
        }
        }
        fmt.Println("]")
    
        }
       
       
   

