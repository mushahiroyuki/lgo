package main

import(
	"fmt"
	"reflect"
)


func main() {

	// 1. まず reflect.TypeOf()の値を得る（reflection的型情報を得る）
	// 2. その型に定義されているメソッドを使う（Name、Kindなど）

	type Foo struct { //liststart1
		A int
		B string
	}

	var x int 
	xt := reflect.TypeOf(x) // xtはxの型に関する情報をもつ
	fmt.Println(xt.Name()) // int  // 型の名前は「int」（文字列）
	fmt.Printf("%T\n", xt.Name()) // string  // xt.Name()の型はstring（%Tで型を表示）

	f := Foo{}
	ft := reflect.TypeOf(f) // ftはfの型に関する情報をもつ
	fmt.Println(ft.Name())  // Foo // 型の名前は「Foo」（文字列）
	fmt.Printf("%T\n", ft.Name()) // string // ft.Name()の型はstring

	xpt := reflect.TypeOf(&x) // xptはxのポインタの型に関する情報をもつ
	fmt.Println(xpt.Name()) // ""（何も表示されない）// 型の名前はなし
	fmt.Printf("%T\n", xpt.Name()) // string  //listend1

	fmt.Println("-----")
	fmt.Println(xt.Kind()) // int //liststart2
	fmt.Println(ft.Kind()) // struct
	fmt.Println(xpt.Kind()) // ptr //listend2

	fmt.Println("-----")	
	fmt.Printf("%T（%d）\n", xt.Kind(), xt.Kind()) // reflect.Kind(2) // iotaを使って定義された定数 具体的な値は2
	fmt.Printf("%T(%d)\n", ft.Kind(), ft.Kind()) // reflect.Kind(25)
	fmt.Printf("%T(%d)\n", xpt.Kind(), xpt.Kind()) // reflect.Kind(22)
	// reflext.Kindの定義は https://pkg.go.dev/reflect#Kind にあり
	
}
