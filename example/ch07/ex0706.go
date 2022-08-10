package main

import (
	"fmt"
)

type Score int //liststart1
type HighScore Score

type Person struct { // 人
	LastName string  // 姓
	FirstName string // 名
	Age int  // 年齢
} 
type Employee Person // 従業員 //listend1


func (s Score) Double() Score {  //liststart3
	return s * 2
}  //listend3

func main() {
	// 型のない定数の代入は認められる //liststart2
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	// hs = s    // コンパイル時のエラー！
	// s = i     // コンパイル時のエラー！
	s = Score(i)   // 型変換後に代入
	hs = HighScore(s)  // 型変換後に代入
	fmt.Println(s, hs) // 300 300
	hhs := hs + 20  //基底型（int）に対して使える演算子（+）は使える
	fmt.Println(hhs)  // 320  //listend2

	s = 200   //liststart4
	hs = 300
	fmt.Println(s.Double()) // 400
	fmt.Println(Score(hs).Double()) // 600
	// fmt.Println(hs.Double()) // コンパイル時のエラー //listend4
} 
