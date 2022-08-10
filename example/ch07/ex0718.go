package main

import (
	"fmt"
	"io"
	"os"
)

type MyInt int

func main() {
	var i any
	doTypeSwitch(i)
	
	var mine MyInt = 20
	i = mine
	doTypeSwitch(i)

	s := "これは文字列"
	doTypeSwitch(s)

	s2 := []rune(s)
	doTypeSwitch(s2)

	doTypeSwitch(s2[0])

	b := int(mine) < 20
	doTypeSwitch(b)

	b = int(mine) == 20
	doTypeSwitch(b)
	
	type Person struct {
		FirstName string
		LastName string
		Age int
	}
	
	st := Person {
		FirstName: "John",
		LastName: "Baker",
		Age: 20,
	}
	doTypeSwitch(st)


	f, err := os.Open("ex0701.go")

	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	doTypeSwitch(f)

	finalInt := 100
	doTypeSwitch(finalInt)
}


func doTypeSwitch(i any) { //liststart1
	switch j := i.(type) {
	case nil:
		fmt.Printf("case nil; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case int:
		fmt.Printf("case int; i:%d（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case MyInt:
		fmt.Printf("case MyInt; i:%d（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case io.Reader:
		fmt.Printf("case io.Reader; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case string:
		fmt.Printf("case string; i:%s（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case bool, rune:
		fmt.Printf("case bool, rune; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	default:
		fmt.Printf("default; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	} //listend1
}
