package main

import(
	"fmt"
	"reflect"
)

func main() {
	i := 10  //liststart1
	iv := reflect.ValueOf(&i) //listend1
	ivv := iv.Elem() //liststart2
	fmt.Println(ivv) // 10  //listend2
	ivv.SetInt(20) //liststart3
	fmt.Println(ivv) // 20  //listend3
}
