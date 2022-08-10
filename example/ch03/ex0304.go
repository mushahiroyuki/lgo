package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} //liststart
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)  // x: [1 2 3 4]
	fmt.Println("y:", y)  // y: [1 2]
	fmt.Println("z:", z)  // z: [2 3 4]
	fmt.Println("d:", d)  // d: [2 3]
	fmt.Println("e:", e)  // e: [1 2 3 4] //listend
}
