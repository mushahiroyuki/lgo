package main

import "fmt"

func main() {
	m := map[string]int{ //liststart
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("ループ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	} //listend
}
