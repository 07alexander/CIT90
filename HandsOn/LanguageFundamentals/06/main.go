package main

import "fmt"

type person struct{
	fName string
	lName string
	favFood []string
}
func main(){
	p1:= person{"Alexander", "Phosykeo",[]string{"Pho","Spaghetti","Tacos"}}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)
	for i, j := range p1.favFood{fmt.Println(i,j)}
	alex:= p1.walk()
	fmt.Println(alex)
}

func (a person) walk() string{
	return fmt.Sprintln(a.fName, "is walking.")
}
