package main

import (
	"fmt"
	"time"
)

func countTo(max int) <-chan int { //intのチャネルを戻す
	ch := make(chan int)
	go func() {  // 無名関数
		for i := 0; i < max; i++ {
			ch <- i // 0からmax-1までの値を順にchに入れる
		}
		close(ch)
	}()

	return ch // チャネルchを戻す
}

func main() { //liststart
	for i := range countTo(10) { //countToからチャネルが戻り、0...9が代入される 
		fmt.Print(i, " ")
		if i > 5 {
			break
		}
	}
	fmt.Println()
	doSomethingTakingLongTime() // 時間のかかる処理。この間、無名関数は待ち続けている
} //listend

/* 次のようにするのと同じ。上のほうが（慣れれば）簡潔
func main() {
	ch := countTo(10) // int のチャネルが戻る
	for i := range ch { // クローズされるまで読み込まれる
		fmt.Print(i, " ")
	}
	fmt.Println()
}
*/


func doSomethingTakingLongTime() {
	time.Sleep(5*time.Second) // 何か処理をする	
}
