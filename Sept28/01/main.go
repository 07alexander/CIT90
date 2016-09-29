package main

import "fmt"

type square struct{
	name string
	color string
	length int
	width  int
}
type circle struct{
	name string
	color string
	diameter int
}

func main(){
	s1:= square{"A","Black",2, 4}
	c1:= circle{"A","Blue",20}
	fmt.Println(s1)
	fmt.Println(c1)
}
