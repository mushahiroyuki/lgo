package main

import "fmt"

func main() {

// 
	fmt.Println("=====================")
	fmt.Println("[]int{}——intのスライスを要素なしで初期化したもの（空のスライスリテラル）")
	fmt.Println("=====================")

	a := []int{}; // 要素なしで初期化
	fmt.Println(a) // []

	a = append(a, 1)
	fmt.Println(a) // [1]
	a = append(a, 2)
	fmt.Println(a) // [1 2]
	

//------------------	
	type person struct {  // person型の宣言
    name string  // 「フィールド名」はnameで、その型はstring
    age  int
    pet  string
	}
	p := person {
		"山田",  //name の具体的な値
		28,  // age
		"ハムスター", // pet
	}

	fmt.Println("=====================") 
	fmt.Println("map[string]interface{}{}") 
	fmt.Println("=====================") 
	
	//	map[string]interface{}{}——string→interface{}のマップを要素なしで初期化したもの
	// Go 1.18以降では map[string]any{}  と書いても同じ
	// interface{}あるいはanyは任意の値を代入できる「型」の役目をする。
	m := map[string]interface{}{} 
	m["文字"] = "abc"
	m["数字"] = 123
	m["構造体"] = p
	m["構造体の要素"] = p.pet

	fmt.Println(m["文字"]) // abc
	fmt.Println(m["数字"])  // 123
	fmt.Println(m["構造体"]) // {山田 28 ハムスター}
	fmt.Println(m["構造体の要素"]) // ハムスター

	{
		// interface{}の代わりにanyを使っても同じ（anyはGo 1.18より導入）
		m := map[string]any{} 
		m["文字"] = "abc"
		m["数字"] = 123
		m["構造体"] = p
		m["構造体の要素"] = p.pet

		fmt.Println(m["文字"]) // abc
		fmt.Println(m["数字"])  // 123
		fmt.Println(m["構造体"]) // {山田 28 ハムスター}
		fmt.Println(m["構造体の要素"]) // ハムスター
	}

}
	
