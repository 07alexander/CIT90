package main

import "fmt"

type vehicle struct{
	doors int
	color string
}
type truck struct{
	vehicle
	fourWheel bool
}
type sedan struct{
	vehicle
	luxury bool
}

func main(){
	datsun:= truck{
		vehicle{
			2,
			"Black",
		},
		true,
	}
	fmt.Println(datsun)
	fmt.Println(datsun.color)
	fmt.Println(datsun.doors)
	fmt.Println(datsun.fourWheel)
}