package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)

	go func () {
		v := 1
		fmt.Printf("この下でch1へ%vを入れる\n", v)
		ch1 <- v
	}()

	go func () {
		v := 2
		fmt.Printf("この下でch2へ%vを入れる\n", v)
		ch2 <- v
	}()

	go func () {
		fmt.Printf("この下でch3から値を受け取る\n")
		v := <-ch3
		fmt.Printf("ch3から%vを受け取った\n", v)
	}()

	go func () {
		v := 4
		fmt.Printf("この下でch4へ%vを入れる\n", v)		
		ch4 <- v
		fmt.Printf("ch4へ%vを入れた\n", v)
	}()

	x := 3

	select { //liststart
	case v := <-ch1:
		fmt.Println("ch1:", v)
	case v := <-ch2:
		fmt.Println("ch2:", v)
	case ch3 <- x:
		fmt.Println("ch3へ書き込み: ", x)
	case <-ch4:
		fmt.Println("ch4から値をもらったが、値は無視した")
	} //listend
}
