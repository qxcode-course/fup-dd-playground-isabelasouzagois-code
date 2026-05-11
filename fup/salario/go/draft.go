package main
import "fmt"
func main() {
    var salario float32
    fmt.Scan(&salario)
    if salario <= 1000{
        salario := salario+salario*0.20
        fmt.Printf("%.2f\n",salario)
    }else if salario <= 1500{
        salario := salario+salario*0.15
        fmt.Printf("%.2f\n",salario)
    }else if salario <= 2000{
        salario := salario+salario*0.10
        fmt.Printf("%.2f\n",salario)
    }else if salario > 2000{
        salario := salario+salario*0.05
        fmt.Printf("%.2f\n",salario)
    }
    
}
