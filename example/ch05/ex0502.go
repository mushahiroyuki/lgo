package main

import "fmt"

type MyFuncOpts struct { // MyFuncのオプション引数
	FirstName string
	LastName string
	Age int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	fmt.Println("【ここで必要な処理を行う】")
	return nil
}

func main() {
	MyFunc(MyFuncOpts {
		LastName: "Patel",
		Age: 50,
	})
	MyFunc(MyFuncOpts {
		FirstName: "Joe",
		LastName: "Smith",
	})
}
