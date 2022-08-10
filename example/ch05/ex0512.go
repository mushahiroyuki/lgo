package main

import "fmt"

func makeMult(base int) func(int) int { //liststart1
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(2)  // 2倍する関数
	threeBase := makeMult(3)  // 3倍する関数
	for i := 0; i <= 5; i++ {
		fmt.Print(i, ": ", twoBase(i), ", ", threeBase(i), "\n")
	}
} //listend1
