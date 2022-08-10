package main

import (
	"fmt"
)

func main() {
//liststart	
	intMap := map[string]int{} // 文字列→整数のマップで要素がないもの 
	intMap["abc"] = 1
	intMap["あいう"] = 10
	fmt.Println(intMap["abc"]) // 1
	fmt.Println(intMap["あいう"]) // 10

	m := map[string]any{  // interface{}も可 （go 1.18でanyが導入された）
		// any にはどのような型の値も入れられる
		"文字": "あいう",
		"数字": 123,
		"マップ": map[string]int {
			"加奈": 32,
			"涼子": 23,
		},
	}
	fmt.Println(m["文字"]) // あいう
	fmt.Println(m["数字"]) // 123
	fmt.Println(m["マップ"]) // map[加奈:32 涼子:23]
	x := m["マップ"].(map[string]int) // 「7.10.1  型アサーション」
	// m["マップ"]のままではanyなので、そのまま m["マップ"]["加奈"]とはできない
	// 型アサーションで、map[string]int だと（強制的に）仮定する
	fmt.Println(x["加奈"]) // 32 
	fmt.Printf("%T\n", x) // map[string]int // %Tで型を出力
//listend

	fmt.Println("==============")

	l := map[string]map[string]int{
		"マップ1": map[string]int {
			"加奈": 32,
			"涼子": 23,
		},
		"マップ2": map[string]int {
			"加奈": 178,
			"涼子": 158,
		},
	}

	fmt.Println(l["マップ1"]["加奈"]) // 32
	fmt.Println(l["マップ2"]["涼子"]) // 158
}
