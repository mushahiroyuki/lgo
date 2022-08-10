package main

import "fmt"

// 型Personを定義   //liststart1
type Person struct {
	LastName string  // 姓
	FirstName string // 名
	Age int  // 年齢
}  //listend1

// 型Personに付随するメソッドStringを定義（PersonにメソッドStringを付加） //liststart2
func (p Person) String() string { // 「(p Person)」がレシーバの指定
	return fmt.Sprintf("%s %s：年齢%d歳", p.LastName, p.FirstName, p.Age)
} //listend2

func main() { //liststart3
	p := Person { // Person型の変数pの宣言と初期化
		LastName: "武田",
		FirstName: "信玄",
		Age: 52,
	}
	
	output := p.String() // Personに付随するメソッドStringを呼び出す
	fmt.Println(output) // 武田 信玄：年齢52歳
} //listend3
