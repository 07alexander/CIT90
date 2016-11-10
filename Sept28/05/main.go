package main

import (
	"net"
	"log"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil{
		log.Println(err)
	}
	io.WriteString(conn, "hello from callin in")

}
