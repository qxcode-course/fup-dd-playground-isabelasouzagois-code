package main
import "fmt"
func main() {
    var a,b int
    fmt.Scan(&a,&b)
    for i:=a;i<=b;i++{
        if i%3==0&&i%5==0{
            fmt.Println("zigzag")
        }else if i%3==0{
             fmt.Println("zig")
        }else if i%5==0{
             fmt.Println("zag")
        }else{
            fmt.Println(i)
        }
        
    }
    
}
