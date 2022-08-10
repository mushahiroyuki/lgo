package main

import "fmt"

func main() {
	s := "文字列"
	fmt.Printf("%s\n", s) // 「文字列」
	s = `改行や
	タブが入った
		文字列`
	fmt.Printf("%s\n", s) // そのまま表示
	fmt.Printf("%q\n", s) // エスケープされる 

	i := 254
	fmt.Printf("%d %b %o %x %X\n", i, i, i, i, i) // 254 11111110 376 fe FE

	

	fmt.Printf("%#v\n", s) // Goのリテラル表現。文字列の場合 %q と同じ

	


}
