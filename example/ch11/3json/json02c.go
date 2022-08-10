package main

import (
	"encoding/json"
	"time"
	"os"
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
	file, err := os.Open("testdata/data.json")
	            // func Open(name string) (*File, error)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var o Order
	d := json.NewDecoder(file) // 下記コメント参照
	d.Decode(&o)
	fmt.Printf("%+v\n", o)
}

/*
	 func NewDecoder(r io.Reader) *Decoder // パッケージencoding/json
   Fileは io.Readerを満たす。下が定義されている // パッケージos
       func (f *File) Read(b []byte) (n int, err error)

   io.Readerの定義は次のとおり
   type Reader interface {  // パッケージio
       Read(p []byte) (n int, err error)
   }
  */
