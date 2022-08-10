package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch := make(chan int, len(a))
	for _, v := range a { //liststart
		v := v  // 外側のvをシャドーする
		go func() {
			ch <- v * 2
		}()
	} //listend
	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
