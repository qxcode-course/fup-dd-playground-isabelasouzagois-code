package main
import "fmt"
func main() {
    var pedra int
    fmt.Scan(&pedra)
    fmt.Print("[ ")
    for i := 0; i <= 10; i++{

        if i == pedra{
            continue
        }

        if i == 10{
            fmt.Printf("ceu")
            break
        }
        fmt.Printf("%d", i)
        if i < 10{
            fmt.Printf(" ")
        }
    }
    fmt.Printf(" ]")
}
