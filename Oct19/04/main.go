package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request){
	fmt.Println(req.URL.Path)
	fmt.Println(w, "Check Terminal")
}
