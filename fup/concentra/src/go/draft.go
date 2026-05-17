package main
import "fmt"
func main() {
    var a,b int
    fmt.Scan(&a,&b)
    fmt.Print("[ ")
    for i:=a;a<=b;i++{
        fmt.Printf("%d %d ",i,b)
        b--
    }
    fmt.Println("]")
}
