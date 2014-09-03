package main

import "github.com/go-martini/martini"
import "net/http"
import "log"
func main(){
	m := martini.Classic()
	m.Get("/",func() string{
	  return "Hello world!"
	})
	//m.Run(":8989")
	log.Fatal(http.ListenAndServe(":3000",m))
}
