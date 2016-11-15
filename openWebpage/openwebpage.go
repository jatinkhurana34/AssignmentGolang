

    // http_get101.go
    //
    // get the contents of a web page with given URL
    //
    // for imported package info see ...
    // http://golang.org/pkg/fmt/
    // http://golang.org/pkg/io/ioutil/
    // http://golang.org/pkg/net/http/
    //
    // tested with Go version 1.4.2   by vegaseat  28apr2015
    package main
    import (
    	"fmt"
        "openWebpage/openpage"
"encoding/json"
    
    )

func Printformatstruct(lang openpage.Lang) {

    fmt.Printf("%v : \n", lang)

}

func Printjsonstruct(lang openpage.Lang) {

    bolb,_ := json.Marshal(lang)
    fmt.Println(string(bolb))
}



    func main() {



              pythonLang := openpage.Lang{
                             Name:      "Python",
                             Url:       "https://www.python.org/",

              }

              rubyLang := openpage.Lang{
                             Name:      "ruby",
                             Url:       "https://www.ruby-lang.org/en/",

              }

              goLang := openpage.Lang{
                             Name:      "go",
                             Url:       "https://golang.org/",

              }

  openpage.Crawl(Printformatstruct,pythonLang)
  openpage.Crawl(Printjsonstruct,pythonLang)

  openpage.Crawl(Printformatstruct,rubyLang)
  openpage.Crawl(Printjsonstruct,rubyLang)

  openpage.Crawl(Printformatstruct,goLang)
  openpage.Crawl(Printjsonstruct,goLang)



        
    	

    }
