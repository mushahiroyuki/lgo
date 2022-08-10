package main

import "fmt"

func main() { //liststart
	samples := []string{"hello", "apple_π!", "これは漢字文字列"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' || r=='は' {
				continue outer
			}
		}
		fmt.Println()
	}
} //listend
