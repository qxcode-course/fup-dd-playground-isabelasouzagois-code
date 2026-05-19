 package main
import "fmt"
func main() {
    var p,s,e int
    fmt.Scan(&p,&s,&e)
    sapo:=0
    posin:=0
    fmt.Printf("%d",posin)
    for{
        sapo=sapo+s
        if sapo>=p{
             fmt.Println(" saiu")
             break
        }
        fmt.Println("",sapo)
        sapo =sapo -e
        fmt.Printf("%d",sapo)
    }
   
}
