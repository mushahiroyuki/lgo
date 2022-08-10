// 4章のリスト番号の付いていないコードを試すためのものです。
// エラーになる例は、コメントになっています。

package main

import "fmt"

func main() {
	fmt.Println("----- 4.3.4　breakとcontinue -----")

	{
		x := 0
		for	{
			fmt.Println("x: ", x)
			x++
	    if 0<=x {
				break
			}
		}
	}



}

