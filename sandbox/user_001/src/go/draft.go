package main
import "fmt"
func ler (arr []int,sep string){
    fmt.Print("[")
    for i, valor := range arr{
        if i !=0{
            fmt.Print(sep)
        }
        fmt.Printf("%v" , valor)
    }
    fmt.Print("]\n")

}
func main() {

    var qnt int
    fmt.Scan(&qnt)
    var arr [] int = make([]int, qnt)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
   ler(arr ," ")
}
