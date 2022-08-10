package main

import "fmt"

func main() {
	ch := make(chan string)
	go sub(ch) // subをchを引数としてゴルーチンとして起動する

	s := <-ch
	fmt.Println("sをchから読み込み:", s) // sをchから読み込み: メッセージ1
	<-ch // スキップ。2つ目を飛ばすことになる
	s = <-ch	
	fmt.Println("sをchから読み込み:", s) // sをchから読み込み: メッセージ3
}

func sub(c3 chan<-string) { // チャネルに書き込みのみを行う指定
	s1 := "メッセージ1"
	fmt.Println("c3へ書き込み:", s1) // c3へ書き込み: メッセージ1
	c3 <- s1
	s1 = "メッセージ2"
	fmt.Println("c3へ書き込み:", s1) // c3へ書き込み: メッセージ2
	c3 <- s1
	s1 = "メッセージ3"
	fmt.Println("c3へ書き込み:", s1) // c3へ書き込み: メッセージ3
	c3 <- s1
}
