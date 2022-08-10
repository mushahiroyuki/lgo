package main


import (
	"fmt"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
} 

func plusOneThousand[T Integer](in T) T { //liststart1
	return in + 1_000  // _は桁の区切り（「2.1.2 リテラル」参照）
} //listend1


func main() {
	var a int = 10
	b := plusOneThousand(a)

	fmt.Printf("%T: %v\n", b, b)
}
