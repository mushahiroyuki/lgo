package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12} // 偶数
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
}
