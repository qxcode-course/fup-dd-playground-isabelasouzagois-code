package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
vetor := make([]int, num)
      if(num == 0){
        fmt.Print("[]\n")
        
    }  else{
	for i := 0; i < num; i++ {
		fmt.Scan(&vetor[i])
	}
    fmt.Printf("[")

	for i := 0; i < num; i++ {
		if i == num -1 {
			fmt.Print(vetor[i] )
		}else{
            fmt.Print(vetor[i],", ")
        }
        
			
		}
      
        fmt.Printf("]\n")
    }
	}

