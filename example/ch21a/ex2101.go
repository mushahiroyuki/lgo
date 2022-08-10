package main

import "fmt"

func main() {
	s := "隣の客はよく柿食う客だ"
	for _, r := range s {
		fmt.Printf("%s\n", string(r))
	}
}
	
