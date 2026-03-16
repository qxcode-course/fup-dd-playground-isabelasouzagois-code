package main

import "fmt"

func main (){
	var a,b int
  fmt.Scan(&a)
  fmt.Scan(&b)
  var div int = a/b
  var resto int = a%b
  fmt.Println(div,resto)

}
 