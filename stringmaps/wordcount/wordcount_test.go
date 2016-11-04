package wordcount

import ( "testing"
)

func TestFrequency(t *testing.T) {
val := Frequency("go bat cat gopher lang go toy go")

    
    if val["go"] != 3 {
        t.Error("Found error ")
        return
    }  
    if val["bat"] != 1 {
        t.Error("Found error")
        return
    }  
    if val["cat"] != 1 {
        t.Error("Found error")
        return
    }  
    if val["gopher"] != 1 {
        t.Error("Found error")
        return
    }  
    if val["lang"] != 1 {
        t.Error("Found error")
        return
    }  
    if val["toy"] != 1 {
        t.Error("Found error")
        return
    }  

}
