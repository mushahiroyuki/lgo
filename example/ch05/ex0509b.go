package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

type opFuncType func(int,int) int  //liststart1
//listend1

var opMap = map[string]opFuncType{ //liststart2
	"+": add,  // 関数定義の中身は前と同じ //listend2
	"-": sub,
	"*": mul,
	"/": div,  
} //liststart3
//listend3

func main() {
	expressions := [][]string{ //liststart3
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 { // 演算子と被演算子の合計個数のチェック
			fmt.Print(expression , " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0]) // 1番目の被演算子のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]   // 演算子のチェック
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])  // 2番目の被演算子のチェック
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result := opFunc(p1, p2) // 実際の計算
		fmt.Print(expression, " → ", result, "\n")
	} //listend3
}
