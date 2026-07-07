package main
import "fmt"
type pessoa struct{
nome string
idade int
sexo string}
func main() {
    var n int
    fmt.Scan(&n)
    pessoas:=make([]pessoa,n)
    for i:=0;i<n;i++{
        fmt.Scan(&pessoas[i].nome,&pessoas[i].idade,&pessoas[i].sexo)
    }
    maioridade:=-1
    nome := ""
    for i:=0;i<n;i++{
        if pessoas[i].sexo=="f" &&pessoas[i].idade>maioridade{
            maioridade=pessoas[i].idade
            nome=pessoas[i].nome
        }
    }
    if maioridade==-1{
        fmt.Println("nao tem mulher")
    }else{
         fmt.Println(nome)
        
    }
    
}