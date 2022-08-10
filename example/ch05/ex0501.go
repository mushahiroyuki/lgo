package main

import "fmt"

func div(numerator int, denominator int) int { // 分子, 分母
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func main() {
	result := div(5,2)
	fmt.Println(result)
}
