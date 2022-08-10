package main

import "fmt"

func main() {
	x := 10  //liststart1
	pointerToX := &x
	fmt.Println(pointerToX) // アドレスが表示される
	fmt.Println(*pointerToX) // 10  // デリファレンスする
	z := 5 + *pointerToX
	fmt.Println(z) // 15   //listend1

	var y *int  //liststart2
	fmt.Println(y == nil) // true
	// fmt.Println(*y)  // ←パニックになる  //listend2

	
	var a = new(int)  //liststart3
	fmt.Println(a == nil) // false
	fmt.Println(*a)   // 0  //listend3
}
