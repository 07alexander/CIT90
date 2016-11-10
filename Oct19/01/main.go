package main

import (
	"net/http"
	"io"
)

func main(){
	http.HandleFunc("/", idx)
	http.HandleFunc("/picture.png",pic)
	http.ListenAndServe(":8080",nil)
}

func idx(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/picture.png">`)
}
func pic(w http.ResponseWriter, req *http.Request){
	http.ServeFile(w, req, "http://www.unixstickers.com/image/cache/data/stickers/golang/golang.sh-600x600.png")
}
