package main

import "fmt"

func main() {
	var com1, com2, com3 int
	fmt.Scan(&com1, &com2, &com3)
	if com1 == com2 && com2 == com3{
		fmt.Println(3)
	}else if com1 == com2 ||com1==com3||com2 ==com3{
        fmt.Println(2)
    }else {
        fmt.Println(0)
    }
}     

