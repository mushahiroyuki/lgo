
package main

import (
	"fmt"
	"reflect"
)

// reflect.TypeOf(x) で xの型がわかります（14章参照）
func main() {
	x := []int{1, 2, 3, 4}  // スライス
	fmt.Println("x:", x, reflect.TypeOf(x)) // [1 2 3 4]   []int
	d := [4]int{5, 6, 7, 8} // 配列
	fmt.Println("d:", d, reflect.TypeOf(d)) // [5 6 7 8]  [4]int
	y := make([]int, 2)  // スライス
	fmt.Println("y:", y, reflect.TypeOf(y)) // [0 0]  []int
	copy(y, d[:])
	fmt.Println(y) // [5 6]  （yの長さは2なので、2個だけコピー）
	copy(d[:], x)
	fmt.Println("d:", d, reflect.TypeOf(d)) // [1 2 3 4] [4]int
	copy(d[2:3], x)  // [3]が[1]で置き換えられる
	fmt.Println("d:", d) // [1 2 1 4]

	copy(d[2:4], x)  // [3 4] が [1 2]で置き換えられる
	fmt.Println("d:", d) //  [1 2 1 2]

	copy(d[2:], x)  // [3 4] が [1 2]で置き換えられる
	fmt.Println("d:", d) //  [1 2 1 2]

	copy(d[1:], x)  // [2 3 4] が [1 2 3] で置き換えられる
	fmt.Println("d:", d) //  [1 1 2 3]
}
