package calculatefrequency

import(
"fmt"
"strconv"

)

//calculate the frequency and print it
func Frequency( s[]string) ([]int,  string){
    q := make([]int,10)
    for v:= 0; v < len(s); v++{
        //convert the string char to an integer if err is nil then we are good otherwise we will exit by      //showing a message
        value,err := strconv.Atoi(s[v])
	if err !=nil{
             fmt.Println("You entered other than digit.")
             return q,"error"
	} else{
		q[value]++
	}
        
    }
    print(q)
return q,"ok"
}

