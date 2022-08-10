package main

import "fmt"

func main() { //liststart
	a := 10
	goto skip
	b := 20
skip:
	c := 30
	fmt.Println(a, b, c)
	if c > a {
		goto inner
	}
	if a < b {
	inner:
		fmt.Println("aはbより小さい")
	}
} //listend
