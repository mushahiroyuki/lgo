package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i4 := i.(int) // iに代入された型はMyIntなので、パニック！ //liststart
	fmt.Println(i4) //listend
}
