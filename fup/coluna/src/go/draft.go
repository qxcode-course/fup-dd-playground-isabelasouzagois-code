package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    var matriz[100][100]int

    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            fmt.Scan(&matriz[i][j])
        }
    }

        maiorColuna := 0
        maiorValor := 0

        for j := 0; j < n; j++{
            soma := 0

            for i := 0; i < n; i++{
                soma += matriz[i][j] * matriz[i][j]
            }

            if j == 0 || soma > maiorValor{
                maiorValor = soma
                maiorColuna = j
            }
        }
            fmt.Println(maiorColuna)
}