package main

import(
	"fmt"
	"reflect"
)

func main() {
	var x int //liststart1
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())  // "" （空文字列）
	fmt.Println(xpt.Kind())  // ptr
	fmt.Println(xpt.Elem().Name()) // int  （型を表す文字列）
	fmt.Println(xpt.Elem().Kind()) // int  （reflect.Kindの定数）
//listend1

	s := xpt.Elem().Name()
	fmt.Println(reflect.TypeOf(s)) // string  （型を表す文字列）
	c := xpt.Elem().Kind()
	fmt.Println(reflect.TypeOf(c)) // reflect.Kind
}
