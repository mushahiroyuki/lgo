package main

import (
	"fmt"
)

type MyInt int //liststart

func main() {
	var i any
	var mine MyInt = 20
	i = mine
	i2 := i.(MyInt) // 型アサーション。iをMyInt型だと仮定（assert）して値（20）をもらう
	fmt.Println(i2) // 20
}
