package wordcount

import(
"strings"
)

func Frequency(input string) (map[string]int) {
         
        //To split the string with spaces  
        s := strings.Fields(input)

        //create a map with key as string and int as an value
        q := make(map[string]int)

        for v:= 0;v<len(s);v++ {
        q[strings.ToLower(s[v])]++
    }

return q
}
