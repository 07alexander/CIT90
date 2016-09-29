package main

import "fmt"

func main(){
	m:= map[string]int{"Alex":0, "Amanda":1, "Kombu":2}
	fmt.Println(m)
	for i:= range m{fmt.Println(i)}
	for i,j:= range m{fmt.Println(i,j)}
}