package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/",pic)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080",nil)
}

func pic(w http.ResponseWriter, _ *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/resources/gopherswrench.jpg">`)
}
