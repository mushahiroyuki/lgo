package main

import "fmt"

type person struct { //liststart1
	age  int
	name string
} //listend1

func modifyFails(i int, s string, p person) { //liststart2
	i *= 2
	s = "さようなら"
	p.name = "Bob"
} //listend2

func main() { //liststart3
	p := person{}
	i := 2
	s := "こんにちは"
	fmt.Println(i, s, p)  // 2 こんにちは {0 }	
	modifyFails(i, s, p)
	fmt.Println(i, s, p)  // 2 こんにちは {0 }
} //listend3
