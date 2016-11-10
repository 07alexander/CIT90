package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args[1])>=1{
		fmt.Println(os.Args[1])
	}else{
		fmt.Println("HEllo, I am Gopher")
	}
}

