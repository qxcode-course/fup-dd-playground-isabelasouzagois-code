  package main
import "fmt"
func main() {
    var id int
    fmt.Scan(&id)
    num :=id
    inverso:=0
    for{
        ud:=id %10
        inverso=inverso*10+ud
        id=id/10
        if id==0{
            break
        }
    }
    if num==inverso{
        fmt.Println("1")
    }else{
         fmt.Println("0")
    }
   
}
