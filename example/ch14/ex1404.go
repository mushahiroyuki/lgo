package main

import(
	"fmt"
	"reflect"
)

func main() {
	s := []string{"a", "b", "c"} //liststart1
	sv := reflect.ValueOf(s) // svの型はreflect.Value
	s2 := sv.Interface().([]string) // 型アサーション（7.10.1参照）
	// stringのスライスだと仮定して値をs2に代入 //listend1

	fmt.Println(s) // [a b c]
	fmt.Println(sv) // [a b c]
	fmt.Printf("%T\n", sv)  // reflect.Value  （型）
	fmt.Println(s2) // [a b c]
	fmt.Printf("%T\n", s2)  // []string

	s3 := sv.Interface() // 型アサーション
	fmt.Printf("s3:%T\n", s3) 
	fmt.Printf("s3の型: %T\n", s3)
}
