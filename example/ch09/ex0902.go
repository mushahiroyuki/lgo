package main

import "fmt"

type T1 struct { //liststart1
	x int
	S string
}

func (f T1) Hello() string {
	return "hello"
}

func (f T1) goodbye() string {
	return "goodbye"
} //listend1


type T2 = T1 //liststart2
//listend2

func MakeT2() T2 { //liststart3
	t2 := T2 {
		x: 20,
		S: "Hello",
	}
	var f T2 = t2
	fmt.Println(f.Hello())
	return t2
} //listend3

func main() {
	a :=MakeT2()
	fmt.Println(a)
}
