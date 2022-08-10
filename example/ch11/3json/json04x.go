package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
//liststart1	
	const data = `
		{"name": "フレッド", "age": 40}
		{"name": "メアリ", "age": 21}
		{"name": "パッド", "age": 30}
	` //listend1
	var t struct {  // 型名は指定なし  //liststart2  
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var r *strings.Reader
	r = strings.NewReader(data) // 文字列dataから io.Readerをつくる
	// func NewReader(s string) *Reader
	
	var dec *json.Decoder
	dec = json.NewDecoder(r)

	var b bytes.Buffer
	encoder := json.NewEncoder(&b) // bにエンコードした結果を入れる
	// b は bytes.Buffer なので、メモリ上に書く

	for dec.More() {
		err := dec.Decode(&t) // 変数tの情報をもとにデコード
		if err != nil {
			panic(err)
		} //listend2
		fmt.Printf("t: %v\n", t)
		err = encoder.Encode(t) // tをエンコード bにバイト列で書く
		if err != nil {
			panic(err)
		}
	}
	out := b.String() // 文字列に変換
	fmt.Printf("out:\n%v\n", out)
}
