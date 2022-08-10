package main

import (
	"fmt"
)

type MyInt int

func main() {
	var i any  // any はGo 1.18で導入されたもの。interface{} と同じ意味
	var mine MyInt = 20
	i = mine
	i3 := i.(string) // iをstring型だと仮定して値をもらおうとするが... パニック！   //liststart1
	fmt.Println(i3) //listend1
}
