package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(val int) {
			ch <- val * 2
		}(v)
		time.Sleep(time.Millisecond)
		// sleepすれば、実行の順番が決まる
	}
	for i := 0; i < len(a); i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
