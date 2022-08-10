package main

import "fmt"

func main() { //liststart
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		fmt.Println(i, v)
		if i == len(evenVals)-2 {
			break
		}
	}	
} //listend
