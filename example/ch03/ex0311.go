package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}  //liststart
	y := make([]int, 2)  // 長さ2のスライスを作る
	num := copy(y, x) // 先頭から2要素だけコピー
	fmt.Println(y, num) // [1 2] 2  //listend
}
