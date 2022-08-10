package main

import "fmt"

func main() { //liststart
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "：偶数")
		case i%3 == 0:
			fmt.Println(i, "：3で割り切れるが2では割り切れない")
		case i%7 == 0:
			fmt.Println(i, "：ループ終了...はしない！")
			break
		default:
			fmt.Println(i, "：退屈な数")
		}
	}
} //listend
