package main

import "fmt"

func main() {
	x := make([]int, 0, 5) // []  //liststart
	x = append(x, 1, 2, 3, 4) // [1 2 3 4]
	y := x[:2]  // [1 2]
	z := x[2:]  // [3 4]
	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3
	y = append(y, 30, 40, 50) // [1 2 30 40 50]
	x = append(x, 60) // [1 2 30 40 60]
	z = append(z, 70) // [30 40 70]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)  //listend
}
