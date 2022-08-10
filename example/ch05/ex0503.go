package main

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	fmt.Println(addTo(3))  // []
	fmt.Println(addTo(3, 2)) // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8)) // [5 7 9 11]
	a := []int{4, 3}
	fmt.Println(addTo(3, a...)) // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
	// fmt.Println(addTo(3, a))  ←コンパイル時のエラー
	// fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}))  ←コンパイル時のエラー
}
