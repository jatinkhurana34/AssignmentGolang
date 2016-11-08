package main

import(
"fmt"
"frequency/calculatefrequency"
"strings"
)

func main(){
var input string
fmt.Println("Enter a string which contains only numbers")
fmt.Scanln(&input)


//split the input string
 q,err := calculatefrequency.Frequency(strings.Split(input,""))
    if err == "ok" {
        print(q)
    }
}


//prints the slice 
func print(q []int){
    for v:=0; v< len(q); v++{
        fmt.Println()
        fmt.Println(v ,"occur", q[v],"times")
    }    
}



