package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "3でも5でも割り切れる")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "3で割り切れる")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "5で割り切れる")
			continue
		}
		fmt.Println(i)
	}
}
