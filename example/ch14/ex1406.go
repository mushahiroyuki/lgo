package main

import (
	"fmt"
	"reflect"
)

var stringType = reflect.TypeOf((*string)(nil)).Elem()

var stringSliceType = reflect.TypeOf([]string(nil))

func main() {
	ssv := reflect.MakeSlice(stringSliceType, 0, 10) //liststart1

	sv := reflect.New(stringType).Elem()
	sv.SetString("hello")

	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss) // [hello]   //listend1
}
