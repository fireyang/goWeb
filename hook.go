package main

import(
    "fmt"
    "net/http"
    "bytes"
    "log"
)


func handler(w http.ResponseWriter,r *http.Request){
    fmt.Fprintf(w,"Hi there, I love %s!\n",r.URL.Path[1:])
    body := r.Body
    buf := new(bytes.Buffer)
    buf.ReadFrom(body)
    fmt.Fprintf(w,"%s%n",buf.String())
    log.Println(r.URL)
    log.Println(buf.String())
}

func main(){
    http.HandleFunc("/",handler)
    http.ListenAndServe(":4000",nil)
}
