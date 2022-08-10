package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 5; i++ {
		fmt.Println("ループ", i)
		fmt.Println(m)
	}
}
