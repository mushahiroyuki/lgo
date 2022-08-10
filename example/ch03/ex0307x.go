package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	fmt.Println(x)  // []
	x = append(x, 1, 2, 3, 4)
	fmt.Println(x)   // [1 2 3 4]
	y := x[:2]
	fmt.Println("y:", y)  // [1 2]
	z := x[2:]
	fmt.Println("z:", z)  // [3 4]
	
	fmt.Println("cap: ", cap(x), cap(y), cap(z)) // 5 5 3
	y = append(y, 30, 40, 50) // [1 2 30 40 50]
	fmt.Println("1----")	
	fmt.Println("x:", x)  // [1 2 30 40]
	fmt.Println("y:", y)  // [1 2 30 40 50]
	fmt.Println("z:", z)  // [30 40]
	x = append(x, 60)
	fmt.Println("2----")	
	fmt.Println("x:", x)  // [1 2 30 40 60]
	fmt.Println("y:", y)  // [1 2 30 40 60]
	fmt.Println("z:", z)  // [30 40]
	z = append(z, 70)
	fmt.Println("3----")	
	fmt.Println("x:", x)  // [1 2 30 40 70]
	fmt.Println("y:", y)  // [1 2 30 40 70]
	fmt.Println("z:", z)  // [30 40 70]
}
