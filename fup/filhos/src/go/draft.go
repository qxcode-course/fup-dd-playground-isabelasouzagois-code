package main
import "fmt"
func main() {
    var qnt,id int
    fmt.Scan(&qnt,&id)
    valor := qnt
    for i:= 0;i<id ;i++{
        fmt.Println(valor)
        valor +=2
    }
    
}
