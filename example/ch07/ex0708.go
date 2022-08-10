package main

import (
	"fmt"
)


type Employee struct { // 従業員  //liststart1
	Name string
	ID   string
}

func (e Employee) Description() string { // 従業員に関する記述
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct { // マネージャ
	Employee // 型のみ書く（埋め込みフィールド） NameとIDが加わる
	Reports []Employee  // 部下（報告の対象者）  Employeeのスライス
}

func (m Manager) FindNewEmployees() []Employee {  // 新しい従業員を見つける
	newEmployees := []Employee{ // Employee（従業員）のスライス
		Employee{
			"石田三成",
			"13112",
		},
		Employee{
			"徳川家康",
			"13115",
		},
	}
	return newEmployees
}
//listend1

func main() {
	m := Manager{ //liststart2
		Employee: Employee{
			Name: "豊臣秀吉",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID) // 12345
	fmt.Println(m.Description()) // 豊臣秀吉 (12345)

	m.Reports = m.FindNewEmployees()
	fmt.Println(m.Employee)  // {豊臣秀吉 12345}
	fmt.Println(m.Reports)  // [{石田三成 13112} {徳川家康 13115}]
//listend2
}
