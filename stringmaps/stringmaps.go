package main

import(
"fmt"
"bufio"
"os"

"sort"
"stringmaps/wordcount"

)
//main function
func main(){

fmt.Println("Enter a string")

//To take a input from console
//To ignore the output of a function we use _
in := bufio.NewReader(os.Stdin)
input,_ := in.ReadString('\n')



//call method to calculate the frequency of a word and print it
frequency := wordcount.Frequency(input)
sortAndPrint(frequency)
}



//funtion to sort the maps and print it on sonsole
func sortAndPrint(q map[string]int){
      keys := []string{}
    for key,_ := range q{
        keys = append(keys,key)
    }

    //sort the string slice - may be we have a direct method to sort the maps??
    sort.Strings(keys)
    for k := range keys {
        key := keys[k]
        value := q[key]
        fmt.Println(key,"occur", value,"times")
    }   
}

