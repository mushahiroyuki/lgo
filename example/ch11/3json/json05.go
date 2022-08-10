//  ファイルにある複数の値を、メモリ（バッファ）に書き出す
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"bytes"
	"time"
)

// json02.go と同じ形式
type Order struct { // 注文
	ID            string    `json:"id"`  //注文ID
	DateOrdered   time.Time `json:"date_ordered"` //注文日時
	CustomerID    string    `json:"customer_id"` //顧客ID
	Items         []Item    `json:"items"` //商品
	//フィールド名は大文字で始める！ jsonのパッケージから見えない
}

type Item struct {
	ID   string `json:"id"`  //商品ID
	Name string `json:"name"` //商品名
}


func main() {
	r, err := os.Open("testdata/data2.json")	// io.Readerをつくる
	// このファイルには複数のデータが入っている
	if err != nil {
		log.Fatal(err)
	} 
	
	var dec *json.Decoder
	dec = json.NewDecoder(r)

	var b bytes.Buffer
	encoder := json.NewEncoder(&b) // bにエンコードした結果を入れる
	// b は bytes.Buffer なので、メモリ上に書く
	
	for dec.More() {
		var o Order
		err := dec.Decode(&o) // 変数oの情報をもとにデコード
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("o: %v\n", o)
		err = encoder.Encode(o) // tをエンコード bにバイト列で書く
		if err != nil {
			log.Fatal(err)
		}
	}
	out := b.String() // 文字列に変換
	fmt.Printf("out:\n%v\n", out)
}
