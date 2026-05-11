package main
import "fmt"
func main() {
    var jog1,jog2 string
    fmt.Scan(&jog1,&jog2)
    S:="S"
    R:="R"
    P:="P"
    if jog1 == jog2{
        fmt.Println("empate")

    }else if jog1 == P && jog2 == R||jog1==S&& jog2==P||jog1==R&&jog2==S{
        fmt.Println("jog1")
     }else{
        fmt.Println("jog2")
     }
    
    }
    
    
    
   

