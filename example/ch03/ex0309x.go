package main

import "fmt"

func main() {
	x := [4]int{5, 6, 7, 8} // [5 6 7 8]  //liststart
	y := x[:2]  // [5 6]
	z := x[2:]  // [7 8]
	d := x[:] // 配列 -> スライスへの変換
	x[0] = 10 // [10 6 7 8]
	fmt.Println("x:", x) // [10 6 7 8]
	fmt.Println("y:", y) // [10 6]
	fmt.Println("z:", z) // [7 8]
	fmt.Println("d:", d) // [10 6 7 8]	
}
