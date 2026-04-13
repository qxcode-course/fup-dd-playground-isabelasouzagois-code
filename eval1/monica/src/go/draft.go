package main
import "fmt"
func main() {
    var m,f1,f2 int
    fmt.Scan(&m,&f1,&f2)
    f3 := m - f1 - f2

    if f1>=f2&&f1>=f3 {
    fmt.Println(f1)
        }else if f2>=f1&&f2>=f3 {
        fmt.Println(f2)
        }else{
            fmt.Println(f3)
        }
}
    
        
    
    
    

