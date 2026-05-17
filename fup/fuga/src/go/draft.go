package main
import "fmt"
func main() {
    var H,P,F,D int
    fmt.Scan(&H,&P,&F,&D)
   
    for i:=0;i<16;i++{

     if F==H{
        fmt.Println("S")
        return
     }
      if F==P{
            fmt.Println("N")
            return
      }
      F+=D
        if F > 15{
            F = 0
        }
        if F < 0{
            F = 15
        }
        
       
        }
   
        
        }
        
        

        
        
        
    
   

