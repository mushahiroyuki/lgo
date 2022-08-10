package main

import "fmt"

func main() {
loop:  //liststart1
	for i := 0; i < 10; i++ { //listend1
		switch {
		case i%2 == 0:
			fmt.Println(i, "：偶数")
		case i%3 == 0:
			fmt.Println(i, "：3で割り切れるが2では割り切れない")
		case i%7 == 0: //liststart2
			fmt.Println(i, "：ループ終了！")
			break loop //listend2
		default:
			fmt.Println(i, "：退屈な数")
		}
	}
}
