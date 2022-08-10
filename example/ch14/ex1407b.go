package main

import (
	"fmt"
	"reflect"
)

func hasNoValue(i any) bool { //liststart1
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
} //listend1

func main() {
	var a interface{}
	fmt.Println(a == nil, hasNoValue(a)) // prints true true

	var b *int
	fmt.Println(b == nil, hasNoValue(b)) // prints true true

	var c interface{} = b
	fmt.Println(c == nil, hasNoValue(c)) // prints false true

	var d int
	fmt.Println(hasNoValue(d)) // prints false

	var e interface{} = d
	fmt.Println(e == nil, hasNoValue(e)) // prints false false
}
