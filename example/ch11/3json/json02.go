package main

import (
	"encoding/json"
	"time"
	"os"
	"fmt"
	"log"
)


type Order struct { // 注文  //liststart1
	ID            string    `json:"id"`  //注文ID
	DateOrdered   time.Time `json:"date_ordered"` //注文日時
	CustomerID    string    `json:"customer_id"` //顧客ID
	Items         []Item    `json:"items"` //商品
	//フィールド名は大文字で始める！ jsonのパッケージから見えない
}

type Item struct {
	ID   string `json:"id"`  //商品ID
	Name string `json:"name"` //商品名
} //listend1


func main() {
	data := readData()
	var o Order //liststart2
	err := json.Unmarshal(data, &o)
	if err != nil {
		log.Fatal(err)
	} //listend2

	fmt.Printf("%+v\n", o) // %+v でフィールド名付きで出力  //listend2
}

func readData() []byte {
	b, err := os.ReadFile("testdata/data.json")
	// ファイルの内容をバイト列（[]byte）としてすべて読み込み
	if err != nil {
		log.Fatal(err)
	}
	return b
}
