package main

import "fmt"

type gator int
type flamingo bool
type swampCreature interface {
	greeting()
}

func (f gator) greeting() {
	fmt.Println("Hello, I am a gator")
}
func (a flamingo) greeting(){
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}
func bayou(r swampCreature){
	r.greeting()
}

func main() {
	var g1 gator
	bayou(g1)
	var f1 flamingo
	bayou(f1)
}