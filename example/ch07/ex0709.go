package main

import (
	"fmt"
)

type Inner struct { //liststart1
	X int
}

type Outer struct {
	Inner
	X int
} //listend1


func main() {
	o := Outer{ //liststart2
		Inner: Inner{
			X: 10,
		},
		X: 20,
	}
	fmt.Println(o.X)       // 20
	fmt.Println(o.Inner.X) // 10   //listend2
}
