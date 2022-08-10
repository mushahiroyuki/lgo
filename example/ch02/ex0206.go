package main

import "fmt"

func main() { //liststart
	ａ := "こんにちは"  // 全角の「ａ」（Unicode U+FF41）
	a := "サヨウナラ" // 半角の「a」 (Unicode U+0061)
	fmt.Println(ａ)
	fmt.Println(a)
} //listend

