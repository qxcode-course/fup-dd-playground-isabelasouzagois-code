package main
import "fmt"
func main() {
    var num1,num2,num3 int
    fmt.Scan(&num1,&num2,&num3)
    var meio int 
    if (num1> num2 && num1 <num3)||(num1<num2 &&num1>num3) {
        meio =num1
    }else if (num2>num1 && num2<num3)||(num2<num1 && num2>num3){
        meio =num2
    }else{
        meio= num3
    }
    fmt.Println(meio)
}
