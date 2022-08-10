package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct { //liststart1
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println("●初期データ")
	fmt.Println(people) //listend1

	// 姓（LastName）でソート //liststart2
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("●姓（LastName。2番目のフィールド）でソート")
	fmt.Println(people) //listend2

	// 年齢（Age）でソート //liststart3
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("●年齢（Age）でソート")	
	fmt.Println(people) //listend3

	fmt.Println("●ソート後のpeople") //liststart4
	fmt.Println(people) //listend4
}
