package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(Min(a, b))
	type MyInt int
	var myA MyInt = 10
	var myB MyInt = 20
	fmt.Println(Min(myA, myB))

	s := "あいう"
	t := "かきく"
	fmt.Println(Min(s, t))

	s = "山川"
	t = "川本"
	fmt.Println(Min(s, t))

}

type BuiltInOrdered interface { //liststart1
	~string | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
} //listend1

func Min[T BuiltInOrdered](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}
