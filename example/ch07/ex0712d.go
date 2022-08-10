package main

import (
	"fmt"
	"os"   // Go 1.16~
	// "io/ioutil"   // Go 1.15まで
	"encoding/json"
)

func main() {
	getData()
}

func getData() error {
	data := map[string]any{} // string->anyのマップで要素なし //liststart
	contents, err := os.ReadFile("sample.json")
	if err != nil {
		return err
	}
	json.Unmarshal(contents, &data) // これで内容がマップdataに入る
	fmt.Println(data) //listend
	return nil
}
