package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!", "これは漢字文字列"} //liststart
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	} //listend
}
