package main

import (
	"fmt"
	"os"   // Go 1.16~
	// "io/ioutil" // Go 1.15まで
	"encoding/json"
)

func main() {
	getData()
}

func getData() error {
	data := map[string]interface{}{}  // map[string]any{} でも可 //liststart
	// 左側の「{}」は空インタフェース（「interface{}」）のため
	// 右側の「{}」はmap[string]interface{}の空のインスタンスを生成
	
	contents, err := os.ReadFile("sample.json")
	// contents, err := ioutil.ReadFile("sample.json") //こちらでも可だが上が呼ばれる
	if err != nil {
		return err
	}
	json.Unmarshal(contents, &data) //これで内容がマップdataに入る
	fmt.Println(data) //listend
	return nil
}
