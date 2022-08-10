package main

import(
	"fmt"
	"reflect"
)

func main() {
	{
		var i int
		changeInt(&i)
		fmt.Println("i:", i) // i: 20
	}

	{
		var i int
		changeIntReflect(&i)
		fmt.Println("i:", i) // i: 20
	}
}


func changeInt(i *int) { //liststart1
	*i = 20
}

func changeIntReflect(i *int) {
	iv := reflect.ValueOf(i)
	iv.Elem().SetInt(20)
} //listend1
