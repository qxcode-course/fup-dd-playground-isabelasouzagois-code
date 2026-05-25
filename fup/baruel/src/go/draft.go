package main
import "fmt"
func main() {
    var total,qtd int 
    fmt.Scan(&total,&qtd)
    v:=make([]int,qtd)
    freq:=make([]int,total+1)

    for i:=0;i<qtd;i++{
        fmt.Scan(&v[i])
        freq[v[i]]++
    }
    fmt.Print("[ ")
    for i:=1;i<=total;i++{
        for j:=1;j<freq[i];j++{
            fmt.Print(i," ")
        }
    }
    fmt.Println("]")

    fmt.Print("[ ")
    for i:=1;i<=total;i++{
        if freq[i]==0{
            fmt.Print(i," ")
        }
    }
    fmt.Println("]")
}
