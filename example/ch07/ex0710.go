package main

import (
	"fmt"
)

type Employee struct { //liststart0
	Name string
	ID   string
}

type Manager struct {
	Employee
	Reports []Employee
} //listend0

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

func main() {
	m := Manager{ //liststart1
		Employee: Employee{
			Name: "上杉謙信",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	var eOK Employee = m.Employee // OK！
	fmt.Println(eOK)              // {上杉謙信 12345}
	var eFail Employee = m        // コンパイル時のエラー！  //listend1
	fmt.Println(eFail)
}
