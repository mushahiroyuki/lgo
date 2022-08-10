package main

import(
	"fmt"
	"reflect"
)

/*
//liststart1
type Foo struct { // 型（type）の定義
	A int
	B string
}


func DoSomething(f Foo) { // 関数の定義
	fmt.Println(f.A, f.B)
} //listend1
*/

func main() {
	var x int //liststart2
	xt := reflect.TypeOf(x) // xtはreflect.Type型。xの型に関する情報をもつ
	fmt.Printf("xt.Name():%v|\n", xt.Name())  // int
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Printf("ft.Name():%v|\n", ft.Name())  // Foo
	xpt := reflect.TypeOf(&x)
	fmt.Printf("xpt.Name():%v|\n", xpt.Name()) // "" （何も表示されない） //listend2
	
	fmt.Println(xt.Kind()) // int //liststart3
	fmt.Println(ft.Kind()) // struct
	fmt.Println(xpt.Kind()) // ptr //listend3
}
