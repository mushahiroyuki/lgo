package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}  //liststart
	num := copy(x[:3], x[1:]) // [1 2 3]  ← [2 3 4]
	fmt.Println(x, num) // [2 3 4 4] 3  //listend
}
