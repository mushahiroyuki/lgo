package main

import "fmt"


func main() {

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	s := "Perry"
	p := person{
		FirstName:  "Pat",
		MiddleName: &s,
		LastName:   "Peterson",
	}
	fmt.Println(p) // {Pat <アドレス> Peterson}
	fmt.Println(*p.MiddleName) // Perry
}
