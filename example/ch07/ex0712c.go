package main

import "fmt"

func main() {
	var i interface{} // 「var i any」も可  //liststart
	i = 20
	fmt.Println(i) // 20
	i = "hello"
	fmt.Println(i) // hello
	i = struct {
		FirstName string
		LastName string
	} {"信玄", "武田"}
	fmt.Println(i) // {信玄 武田}  //listend

	var i2 any
	i2 = 20
	fmt.Println(i2) // 20
	i2 = "hello"
	fmt.Println(i2) // hello
	i2 = struct {
    FirstName string
    LastName string
	} {"信玄", "武田"}
	fmt.Println(i2) // {信玄 武田}


	type Person struct {
    FirstName string
    LastName string
	}

	var i3 any
	i3 = 20
	fmt.Println(i3) // 20
	i3 = "hello"
	fmt.Println(i3) // hello
	i3 = Person{
    FirstName: "信玄",
    LastName: "武田",
	}
	i4 := i3.(Person)
	fmt.Println(i4.LastName, i4.FirstName) // 武田 信玄
}
