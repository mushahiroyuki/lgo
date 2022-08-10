package main

import "fmt"

func main() {
	var s *string  //liststart
	fmt.Println(s == nil) // true
	var i interface{}  // interface{}の代わりにanyでも同じ
	fmt.Println(i == nil) // true
	i = s
	fmt.Println(i == nil) // false  //listend
}
