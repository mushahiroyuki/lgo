// このファイルはコンパイルできません。
package main

import (
	"fmt"
)

type Stack[T any] struct { // 型Stackの宣言。ジェネリクス版 //liststart1
	vals []T
} //listend1

func (s *Stack[T]) Push(val T) { // メソッドPushの定義
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) { // メソッドPopの定義
	if len(s.vals) == 0 {
		var zero T  // 変数の宣言。zeroはその型のゼロ値になる
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool { //liststart2
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
} //listend2

func main() {
	var intStack Stack[int] 
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println(intStack.Contains(10)) // true
	fmt.Println(intStack.Contains(5)) // false
}
