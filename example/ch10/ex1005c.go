package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch := make(chan int, len(a))
	for _, v := range a { //liststart
		go func(val int) { // valの値は下の引数vの値になる
			ch <- val * 2
		}(v)  // vの値を無名関数に引数で渡す
	} //listend
	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
