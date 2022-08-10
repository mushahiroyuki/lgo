package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4} // オリジナルのスライス //liststart
	y := make([]int, 4)  // 長さ4のスライスy
	fmt.Println("len(y):", len(y), "cap(y):", cap(y)) // 4  4
	num := copy(y, x)  // xからyにコピーする
	fmt.Println(y, num) // [1 2 3 4] 4   //listend
}
