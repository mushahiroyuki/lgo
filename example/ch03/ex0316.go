package main

import "fmt"

func main() {
	totalWins := map[string]int{}  // 合計勝利数 //liststart
	totalWins["ライターズ"] = 1
	totalWins["ナイツ"] = 2
	fmt.Println(totalWins["ライターズ"])  // 1
	fmt.Println(totalWins["ミュージシャンズ"])  // 0
	totalWins["ミュージシャンズ"]++
	fmt.Println(totalWins["ミュージシャンズ"]) // 1
	totalWins["ナイツ"] = 3
	fmt.Println(totalWins["ナイツ"])  // 3  //listend
}
