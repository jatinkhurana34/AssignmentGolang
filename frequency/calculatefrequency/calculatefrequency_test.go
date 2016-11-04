package calculatefrequency

import ( "testing"
"strings"
)

func TestFrequency(t *testing.T) {
val,err := Frequency(strings.Split("849832748700032875432711118578",""))

    if (err == "error") {
        t.Fatalf("Found error : other than number was input")
        return
    }

    if val[0] != 3 {
        t.Error("Found error with number :",0)
        return
    }  
    if val[1] != 4 {
        t.Error("Found error with number :",1)
        return
    } 
    if val[2] != 3 {
        t.Error("Found error with number :",2)
        return
    } 
    if val[3] != 3 {
        t.Error("Found error with number :",3)
        return
    } 
    if val[4] != 3 {
        t.Error("Found error with number :",4)
        return
    } 
    if val[5] != 2 {
        t.Error("Found error with number :",5)
        return
    } 
    if val[6] != 0 {
        t.Error("Found error with number :",6)
        return
    } 
    if val[7] != 5 {
        t.Error("Found error with number :",7)
        return
    }
    if val[8] != 6 {
        t.Error("Found error with number :",8)
        return
    }
    if val[9] != 1 {
        t.Error("Found error with number :",9)
        return
    } 
}
