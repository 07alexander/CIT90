package main

import (
	"io"
	"log"
	"net"
	"fmt"
	"bufio"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()
	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go serve(c) //Adding go in front makes it a go routine
	}
}
func serve(c net.Conn){
	defer c.Close()
	scanner := bufio.NewScanner(c)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprint(c, ln)
		if ln == ""{
			break
		}
	}
	io.WriteString(c, "Hello World")
}

