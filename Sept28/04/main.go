package main

import (
	"net"
	"log"
	"io"
)

func main() {
	listen, err:=net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln()
	}
	conn,err := listen.Accept()
	if err != nil{
		log.Fatalln()
	}
	io.WriteString(conn,"hi")
}
