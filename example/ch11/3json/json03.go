// json03x.goに途中経過の説明あり

package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil" // Go 1.16から ioutil.TempFile は単にos.CreateTempを呼ぶ
	"os"
)

func main() {
	type Person struct { //liststart1
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	dataToFile := Person{ // ファイルにJSON形式で書き込むデータ
		Name: "フレッド",
		Age:  40,
	} //listend1

	// 書き込み
	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-") //liststart2
	
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name()) // 使い終わったら削除
	err = json.NewEncoder(tmpFile).Encode(dataToFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	} //listend2
	fmt.Printf("ファイルに書き込んだデータ: %+v\n", dataToFile) // %vで構造体をフィールド名付きで表示	

	// 読み込み
	tmpFile2, err := os.Open(tmpFile.Name()) //liststart3
	if err != nil {
		panic(err)
	}
	var dataFromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&dataFromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ファイルから読み込んだデータ: %+v\n", dataFromFile) // 構造体をフィールド名付きで表示
//listend3	
}
