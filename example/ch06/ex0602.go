package main

import "fmt"


func main() {

	type person struct { //liststart1
		FirstName  string
		MiddleName *string
		LastName   string
	}
	
	p := person{
		FirstName:  "Pat",
		MiddleName: "Perry",  // ←コンパイル時のエラー
		LastName:   "Peterson",
	}
	fmt.Println(p) //listend1
}

