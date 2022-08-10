package main

import (
	"fmt"
	_	"io"
	"os"
)

type MyInt int

func main() {
	var i interface{}
	useTypeSwitch(i)

	var i2 any
	useTypeSwitch(i2)

	var i3 = 3
	i = i3
	useTypeSwitch(i3)

	var mine MyInt = 20
	i = mine
	useTypeSwitch(i)

	s := "これは文字列"
	useTypeSwitch(s)

	s2 := []rune(s)
	useTypeSwitch(s2)

	useTypeSwitch(s2[0])

	b := int(mine) < 20
	useTypeSwitch(b)

	b = int(mine) == 20
	useTypeSwitch(b)
	
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
	useTypeSwitch(st)


	f, err := os.Open("ex0701.go")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	useTypeSwitch(f)
}


func useTypeSwitch(i interface{}) { //liststart
	switch j := i.(type) {
	case nil:
		fmt.Printf("nilの場合; i:%v（型:%T）  jの型:%T\n", i, i, j)
	case int:
		fmt.Printf("intの場合; i:%d（型:%T）  jの型:%T\n", i, i, j)
	case MyInt:
		fmt.Printf("MyIntの場合; i:%d（型:%T）  jの型:%T\n", i, i, j)
	case *os.File:   /*  io.Reader:  も可 */
		fmt.Printf("*os.Fileの場合; i:%v（型:%T）  jの型:%T\n", i, i, j)
	case string:
		fmt.Printf("stringの場合; i:%s（型:%T）  jの型:%T\n", i, i, j)
	case bool, rune:
		fmt.Printf("boolの場合, rune; i:%v（型:%T）  jの型:%T\n", i, i, j)
	default:
		fmt.Printf("default; i:%v（型:%T）  jの型:%T\n", i, i, j)
	}
} //listend
