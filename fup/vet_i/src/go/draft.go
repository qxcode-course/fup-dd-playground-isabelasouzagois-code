package main
import "fmt"
func vetores (arr []int,sep string){
    
    for i, valor := range arr{
        if i !=0{
            fmt.Println(sep)
        }
        fmt.Printf("%v",valor)
    }
   fmt.Print("\n")
    
}
func main() {

    var qnt int
    fmt.Scan(&qnt)
    var arr [] int = make([]int, qnt)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
   vetores(arr ,"")
}

