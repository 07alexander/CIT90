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
	wrx:= sedan{
		vehicle{
			4,
			"Blue",
		},
		false,
	}
	fmt.Println(datsun.transportationDevice())
	fmt.Println(wrx.transportationDevice())
}

func (a truck) transportationDevice() string{
	return fmt.Sprintln(a.color," Truck with ", a.doors, " doors")
}
func (a sedan) transportationDevice() string{
	return fmt.Sprintln(a.color," Sedan with ", a.doors, " doors")
}