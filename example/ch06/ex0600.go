// 6章の独立したファイルになっていない例です。

package main

import "fmt"

func main() {
	fmt.Println("----- new -----")
	{
		var x = new(int)
		fmt.Println("x == nil:", x == nil) // false
		fmt.Println("*x: ", *x)   // 0
	}

	fmt.Println("------- #4 -------")
	{
		type Foo struct {
			firstName string
			lastName string
		}
		x := &Foo{} // 構造体のポインタ
		x.firstName = "太郎"
		x.lastName = "山田"
		fmt.Println(x.lastName)
		fmt.Println(x.firstName)
		var y string
		fmt.Println(y)
		z := &y
		fmt.Println(z)
	}
	
}

