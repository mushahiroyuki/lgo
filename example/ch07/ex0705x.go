package main

import (
	"fmt"
)

type Adder struct { // 構造体Adderの定義 //liststart1
	start int  // int型のフィールドstartをもつ
}

func (a Adder) AddTo(val int) int { // 型Adderをレシーバとするメソッドを定義
	return a.start + val // フィールドstartの値に、引数の値valを足して戻す
} //listend1

func main() {
	myAdder := Adder{start: 10} // startの値を10にしてAdderのインスタンスを生成 //liststart2
	fmt.Println("myAdder.AddTo(5):", myAdder.AddTo(5)) // 15  //listend2

	f1 := myAdder.AddTo // Adder型の変数myAdderのメソッドAddToをf1に代入 //liststart3
	fmt.Println("f1(10):", f1(10)) // 20  //listend3

	f2 := Adder.AddTo //型Adderをレシーバとして定義されているメソッドAddToをf2に代入 //liststart4
	fmt.Println("f2(myAdder, 15):", f2(myAdder, 15)) // 25 //listend4

} 
