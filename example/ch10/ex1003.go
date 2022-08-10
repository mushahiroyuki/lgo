package main

import "fmt"

func main() { //liststart
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v // ch1に書き込めない限りここで待たされる
		v2 := <-ch2
		fmt.Println(v, v2)
	}()

	v := 2
	ch2 <- v  // ch2に書き込めない限りここで待たされる
	v2 := <-ch1
	fmt.Println(v, v2)
} //listend
