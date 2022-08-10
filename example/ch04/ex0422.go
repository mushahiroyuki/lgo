package main

import (
	"fmt"
	"unicode/utf8" // utf8.RuneCountInString(word)で、wordの文字数を数えられる
)

func main() {
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain",
		"タコの足とイカの足", "antholopologist","タコの足は8本でイカの足は10本"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("「%s」の文字数は%dで、短い単語だ。\n", word, size)
		case 5:
			fmt.Printf("「%s」の文字数は%dで、これはちょうどよい長さだ。\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("「%s」の文字数は%dで、とても長い。", word, size)
			if n := len(word); size < n {
				fmt.Printf("%dバイトもある！\n", n)
			} else {
				fmt.Println()
			}
		}
	}
}
