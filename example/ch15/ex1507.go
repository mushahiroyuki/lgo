package main

import "fmt"

type ImpossiblePrintableInt interface { //liststart1
	int
	String() string
}

type ImpossibleStruct[T ImpossiblePrintableInt] struct {
	val T
}

type MyInt int

func (mi MyInt) String() string {
	return fmt.Sprint(mi)
} //listend1

func main() {
	s := ImpossibleStruct[int]{10} //liststart2
	s2 := ImpossibleStruct[MyInt]{10} //listend2
	fmt.Println(s)
	fmt.Println(s2)
}
