package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20} //liststart
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch, " ")
	} //listend
	fmt.Println()
}
