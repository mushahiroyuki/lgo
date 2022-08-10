package main

import "fmt"

// 無名関数を定義して、その無名関数の外の関数で定義されたchをそのまま使う 
func main() {
	ch := make(chan string) // 文字列をやり取りするチャネルchを作る

	go func () {
		s1 := "メッセージ1"
		fmt.Println("chへ書き込み:", s1) // chへ書き込み: メッセージ1
		ch <- s1  // 外側の関数のチャネル変数chをそのまま使う

		s1 = "メッセージ2"
		fmt.Println("chへ書き込み:", s1) // chへ書き込み: メッセージ2
		ch <- s1

		s1 = "メッセージ3"
		fmt.Println("chへ書き込み:", s1) // chへ書き込み: メッセージ3
		ch <- s1
	}()  // () を書いて無名関数を実行する。()を書かないとコンパイル時のエラー

	s := <-ch
	fmt.Println("sをchから読み込み:", s) // sをchから読み込み: メッセージ1

	<-ch // スキップ。2つ目を飛ばすことになる

	s = <-ch	
	fmt.Println("sをchから読み込み:", s) // sをchから読み込み: メッセージ3  //listend1
}
