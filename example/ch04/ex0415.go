package main

import "fmt"

func main() {
	uniqueNames := map[string]bool{"花子": true, "太郎": true, "洋子": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}
