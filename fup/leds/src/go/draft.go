package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    leds:=[]int{6,2,5,5,4,5,6,3,7,6}
    for i:=0;i<n;i++{
        var v string
        fmt.Scan(&v)

        total:=0

        for _, c:=range v{
            total+=leds[c-'0']
        }
    
    fmt.Printf("%d leds\n", total)
}
}