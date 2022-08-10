package main

import(
	"fmt"
	"reflect"
)

func main() {
	var v int = 12
	vValue := reflect.ValueOf(v) // 

	fmt.Printf("%T(%v)\n", reflect.TypeOf(vValue), vValue)
	fmt.Println(vValue.Type()) // int
	fmt.Println(reflect.vValue.Type()) // int
	fmt.Println(vValue.Kind()) // int
	


	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)        // sv is of type reflect.Value
	s2 := sv.Interface().([]string) // s2 is of type []string
	fmt.Print(s2)

	
}
