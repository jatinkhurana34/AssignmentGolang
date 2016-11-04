package openpage

import (
        "io"
    	"io/ioutil"
    	"net/http"
        "time"
    )

type Lang struct{
Name,Url string
Bytes int64
Timetaken string
}


func Crawl (pfunc func(Lang),lang Lang)  {

        start := time.Now()       

        resp, err := http.Get(lang.Url)
        lang.Timetaken = time.Since(start).String()
    	// handle the error if there is one
    	if err != nil {
    		panic(err)
    	}
    	// do this now so it won't be forgotten
    	defer resp.Body.Close()
    	// reads html as a slice of bytes
    	html, err := io.Copy(ioutil.Discard,resp.Body)

    	if err != nil {
    		panic(err)
    	}
lang.Bytes = html


pfunc(lang)



}
