package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12} //liststart
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals) // [2 4 6 8 10 12]   //listend
}
