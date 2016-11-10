package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/", picture)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080",nil)
}

func picture(w http.ResponseWriter, _*http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src= "/resources/met.gif">`)
}
