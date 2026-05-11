package main
import "fmt"
func main() {
    var n,d,a int
    fmt.Scan(&n,&d,&a)
    if d>=a{
        fmt.Println(d-a)
    }else{
        fmt.Println((n-a)+d)
    }
    
}
