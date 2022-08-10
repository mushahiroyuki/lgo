package main

import(
	"fmt"
	"reflect"
)

type Foo struct {
	A int
	B string
}


func main() {
	f := Foo{} //liststart1
	ft := reflect.TypeOf(f) // ftはfのreflection的型情報
	fmt.Println(ft) // main.Foo
	fmt.Println(ft.Name()) // Foo
	fmt.Println(ft.Kind()) // struct  //listend1
}
