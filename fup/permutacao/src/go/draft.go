package main

import (
	"fmt"
	
)
func invert(s []byte,i,j int){
    for i<j{
        s[i],s[j]=s[j],s[i]
        i++
        j--
    }
}
func proxPermut(s []byte)bool{
    n:=len(s)
    i:=n-2
    for i>=0 && s[i] >=s[i +1]{
        i--
    }
    if i<0{
        return false
    }
    j:=n-1
    for s[j]<=s[i]{
        j--
    }
    s[i],s[j]=s[j],s[i]
    invert(s,i+1,n-1)
    return true
}
func main(){
    var ficha string
    var n int
    fmt.Scan(&ficha,&n)
    s:=[]byte(ficha)
    for i:=0;i<n;i++{
        if !proxPermut(s){
            break
        }
    }
    fmt.Println(string(s))
}