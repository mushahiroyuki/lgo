package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} //liststart
	y := x[:2]
	fmt.Println(cap(x), cap(y))  // 4 4
	y = append(y, 30)
	fmt.Println("x:", x)  // x: [1 2 30 4]
	fmt.Println("y:", y)  // y: [1 2 30]  //listend
}
