package hello

import (
	"appengine"
    "fmt"
    "net/http"
    "encoding/json"
    "log"
)


type Message struct {
    Name string
    Body string
    Time int64
}


func init() {
    http.HandleFunc("/", handler)
}



func handler(w http.ResponseWriter, r *http.Request) {
    
    c := appengine.NewContext(r)
    
    // log url
    c.Infof("Requested URL: %v", r.URL)

    // create message
    m := Message{"Alice", "Hello", 1294706395881547000}

    // marshall message
    b, err := json.Marshal(m)

    if (err != nil) {
    	fmt.Fprint(w, "Unexpected error when marshalling Message")
    	return
    }

    log.Println(string(b))

    fmt.Fprint(w, string(b))
}



