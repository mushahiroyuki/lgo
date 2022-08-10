// os.Openと io.ReadAllを使って読み込み
package main

import (
	"encoding/json"
	"time"
	"os"
	"io"
	"fmt"
	"log"
)


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
	data := readData()
	var o Order
	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", o) // %+v でフィールド名付きで出力
}

func readData() []byte {
	file, err := os.Open("testdata/data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file) // ファイルの内容をすべて読み込み
	// closeもやってくれるので、closeの必要もなし
	// （もっともcloseの対象が返ってこないので、closeしようがない..）
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
