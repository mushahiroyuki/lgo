package main

import (
	"fmt"
)

type Stack[T comparable] struct { // anyをcomparableに変更 //liststart1
	vals []T
} //listend1

func (s *Stack[T]) Push(val T) { // メソッドPushの定義
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() (T, bool) { // メソッドPopの定義
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return top, true
}

func (s Stack[T]) Contains(val T) bool {
	for _, v := range s.vals {
		if v == val {
			return true
		}
	}
	return false
}

func main() { //liststart2
	var s Stack[int] // Stackの要素の型（int）を指定する
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println(s.Contains(10)) // true
	fmt.Println(s.Contains(5)) // false
} //listend2



