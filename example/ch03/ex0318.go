package main

import "fmt"

func main() {
	intSet := map[int]struct{}{} //liststart
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5は含まれている")
	}
	if _, ok := intSet[500]; ok {
		fmt.Println("500は含まれている")
	} //listend
}
