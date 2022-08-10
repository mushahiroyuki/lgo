package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct { //liststart
		Name string // NameのNは大文字！ 小文字だと他パッケージから見えない
		Age int
	}{}
	
	err := json.Unmarshal([]byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`), &f)
	// 大文字小文字の違いを無視して、フィールドに対応付けてくれる
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", f) // {Name:小野小町 Age:20}
	// %+v でフィールド名付きで出力   //listend
}
