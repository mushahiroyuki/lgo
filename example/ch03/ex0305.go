package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} //liststart
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("x:", x)  // x: [10 20 30 4]
	fmt.Println("y:", y)  // y: [10 20]
	fmt.Println("z:", z)  // z: [20 30 4]  //listend
}
