// json01.goに構造体タグを追加した例
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct {
		// 構造体タグを使った例
		Name string `json:"name"`// 構造体タグ（説明は次の例にあり）
		Age int     `json:"age"`
		// NameのN、AgeのAは大文字！ 小文字だと他パッケージで見えない
	}{}
	
	err := json.Unmarshal([]byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`), &f)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(f) // {小野小町 20} 
}
