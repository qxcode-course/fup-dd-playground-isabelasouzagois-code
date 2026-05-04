package main
import "fmt"
func main() {
    var a ,b int
    var soma int = 0
    fmt.Scan(&a,&b)
    if a>b{
        fmt.Println("invalido")
    }else{
    for i:=a;i<=b;i++{
        if i%2==0 && i%3 ==0{
        soma += i
    }
}
    
    fmt.Println(soma)
}
}