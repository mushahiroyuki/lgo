package main

import "fmt"

func countTo(max int) (<-chan int, func()) { //liststart
	ch := make(chan int)
	done := make(chan struct{})
	cancelFunc := func() { // chをクローズする関数を戻す
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancelFunc
}

func main() {
 	ch, cancelFunc := countTo(10)
	for i := range ch {
		if i >= 5 {  // 途中で抜けたくなったら（抜ける条件が整ったら）
			break
		}
		fmt.Print(i, " ")		
	}
	fmt.Println()
	cancelFunc() // chをクローズしてforループを終了する
} //listend
