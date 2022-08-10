package main

import (
	"fmt"
)

// Map は []T1 を[]T2 にマッピング関数fを使って変換する 
// 2つの型パラメータT1、T2をとる
// 任意の型のスライスを扱える
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduceは[]T1を関数を用いて1個の値に「縮約（ひとつの値で代表）」する
// （この場合は、合計をとる）
func Reduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filterは指定された関数を使ってスライスの要素をフィルタリングする
// 関数fがtrueを返した要素だけからなるスライスを返す
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}


func main() {
	words := []string{"One", "Potato", "Two", "Potato", "Three"} //liststart1
	filtered := Filter[string](words, func(s string) bool { // [string]が型パラメータ
		return s != "Potato"
	})
	lengths := Map[string, int](filtered, func(s string) int { // [string, int]
		return len(s)
	})
	
	sum := Reduce[int, int](lengths, 0, func(acc int, val int) int { // [int, int]
		return acc + val
	}) //listend1

	var s = []int{10, 20, 30}
	s = Map(s, func(i int) int { // 自乗する
		return i*i
	})

	fmt.Println(filtered) // [One Two Three]
	fmt.Println(lengths)  // [3 3 5]
	fmt.Println(sum) // 11
	fmt.Println(s) // [100 400 900]
}
