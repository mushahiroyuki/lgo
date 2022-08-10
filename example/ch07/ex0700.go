package main

import (
	"fmt"
)

// iotaと列挙型

type BitField int

const (
	Field1 BitField = 1 << iota // 1が代入される
	Field2                      // 2が代入される
	Field3                      // 4が代入される
	Field4                      // 8が代入される
)


type SomeCategory int

const (
	_ SomeCategory = iota  //このカテゴリは使わない
	CategoryX1
	CategoryX2
	CategoryX3
)


type SomeCategory2 int
const (
	CategoryY1 SomeCategory2 = iota+3
	CategoryY2
	CategoryY3
)


func main() {
	fmt.Println("=== 7.2.7 iotaと列挙型 ===")

	fmt.Println(Field1) // 1
	fmt.Println(Field2) // 2
	fmt.Println(Field3) // 4
	fmt.Println(Field4) // 8

	fmt.Println("---")	
	fmt.Println(CategoryX1) // 1
	fmt.Println(CategoryX2) // 2

	fmt.Println("---")	
	fmt.Println(CategoryY1) // 3
	fmt.Println(CategoryY2) // 4
	
}
