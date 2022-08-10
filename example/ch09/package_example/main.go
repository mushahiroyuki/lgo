package main

import (
	"fmt"
	"example.co.jp/package_example/formatter" // パッケージprint
	"example.co.jp/package_example/math" // パッケージmath
)

func main() {
	num := math.Double(2)  // パッケージmath（math/math.go）
	output := print.Format(num) // パッケージprint（formatter/formatter.go）
	fmt.Println(output)
}
