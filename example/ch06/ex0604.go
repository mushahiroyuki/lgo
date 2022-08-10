package main

import "fmt"


func main() {

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	
	p := person{ //liststart2
		FirstName:  "Pat",
		MiddleName: stringp("Perry"), // これならうまくいく
		LastName:   "Peterson",
	}
	fmt.Println(p) // {Pat 0xc000010250 Peterson}
	fmt.Println(*p.MiddleName) // Perry  //listend2
}

func stringp(s string) *string { //ヘルパー関数 //liststart1
	return &s
} //listend1
