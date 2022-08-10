package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}  //liststart
	d := [4]int{5, 6, 7, 8} // dは配列
	y := make([]int, 2)
	copy(y, d[:])  // ソースは配列dの全要素だが、yの長さは2
	fmt.Println(y) // [5 6]
	copy(d[:], x)  // ターゲットが配列
	fmt.Println(d) // [1 2 3 4]    //listend
}
