package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i interface{}  // any の代わりに interface{} でも同じ
	var mine MyInt = 20
	i = mine
	i3 := i.(string) // iをstring型だと仮定して値をもらおうとするが... パニック！   //liststart1
	fmt.Println(i3) //listend1
}
