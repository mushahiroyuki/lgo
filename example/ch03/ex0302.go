package main

import "fmt"

func main() {
	{
		var data []int  // スライスのゼロ値nilで初期化される（nilスライス）  //liststart1
		fmt.Println(data == nil)  // true    //listend1
	}
	{
		x := []int{} // 空のスライスリテラル //liststart2
		fmt.Println(x == nil)  // false  //listend2
	}
}
