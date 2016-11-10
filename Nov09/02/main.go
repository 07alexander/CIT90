package main

import (
	"net/http"
	"fmt"
	"log"
)

func main(

) {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico",http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, req * http.Request){
	err := req.ParseForm()
	if err!= nil{
		log.Fatalln(err)
	}
	xs:=req.FormValue("q")
	for i, v :=range xs {
		fmt.Fprint(w, "Do search: %v-%v\n ",i, v)
	}
}
