// go run ex22b11.go で3つの例を順に実行します。

package main

import (
	"fmt"
)


// 無名関数に引数でチャネルを指定する
func main() {
	ch := make(chan string)  //liststart2

	go func (c1 chan<-string) { // チャネルに書き込みのみを行う指定
//	go func (c1 <-chan string) {  // ←エラーになる。チャネルから読み込む指定になってしまう
// 	go func (c1 chan<-int) {  // ←エラーになる。書き込む型が違う
		s1 := "メッセージ1"
		fmt.Println("c1へ書き込み:", s1)
		c1 <- s1

		s1 = "メッセージ2"
		fmt.Println("c1へ書き込み:", s1)
		c1 <- s1

		s1 = "メッセージ3"
		fmt.Println("c1へ書き込み:", s1)
		c1 <- s1
	}(ch)

	s := <-ch
	fmt.Println("sをchから読み込み:", s)

	<-ch // スキップ。2つ目を飛ばすことになる

	s = <-ch	
	fmt.Println("sをchから読み込み:", s)  //listend2
}

