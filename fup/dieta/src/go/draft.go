  package main
import "fmt"
func main() {
   var dias int
   fmt.Scan(&dias)

   var calorias = make([]int, dias)

   for i := 0; i < dias; i++ {
    fmt.Scan(&calorias[i])
   }

   soma := 0

   for i := 0; i < dias; i++ {
    soma += calorias[i] 
   }

   media := soma / dias

   fmt.Print(media)
   fmt.Printf(".0\n")

}
