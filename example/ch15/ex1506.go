package main

import (
	"fmt"
)

type BuiltInOrdered interface { //liststart1
	string | int | int8 | int16 | int32 | int64 | float32 | 
		float64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
	// 全体が「型要素（type element）」
	// string、intなどのひとつひとつは「型ターム（type termp）」
} //listend1

func Min[T BuiltInOrdered](v1, v2 T) T { //liststart2
	if v1 < v2 {
		return v1
	}
	return v2
}

func main() {
	a := 10
	b := 20
	fmt.Println(Min(a, b)) //listend2
} 
