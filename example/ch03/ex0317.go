package main

import "fmt"

func main() {
	intSet := map[int]bool{} //liststart
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet)) // 11 8
	fmt.Println(intSet[5])  // true
	fmt.Println(intSet[500])  // false
	if intSet[100] {
		fmt.Println("100は含まれている")
	}
	if intSet[10] {
		fmt.Println("10は含まれている")
	} //listend
}
