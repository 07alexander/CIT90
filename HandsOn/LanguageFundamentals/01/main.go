package main

import (
"fmt"
)

func main() {
	s := []int{0,1,2,3,4,5}
	fmt.Println(s)
	for i := range s {
		fmt.Println(i)
	}
	for i, j := range s {
		fmt.Println(i, j)
	}
}