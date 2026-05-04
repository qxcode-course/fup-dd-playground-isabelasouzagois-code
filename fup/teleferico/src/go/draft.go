package main
import "fmt"
func main() {
    var c,a int
    fmt.Scan(&c,&a)
    viagem := a /(c-1)
    if a%(c-1)!=0{
        viagem++
    }

    fmt.Println(viagem)
}
