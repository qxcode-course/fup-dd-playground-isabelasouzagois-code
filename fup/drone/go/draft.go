package main
import "fmt"
func main() {
    var a,b,c,h,l int
    fmt.Scan(&a,&b,&c,&h,&l)
    areaj:=h*l
    if areaj >=a*b||areaj>=a*c||areaj>=b*c{
        fmt.Println("S")
    }else{
        fmt.Println("N")
    }
    
}

