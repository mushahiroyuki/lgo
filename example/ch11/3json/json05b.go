//  ファイルにある複数の値を、別のファイルに書き出す
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	// "bytes"
	"time"
	"os/exec"
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

const (
	inFileName = "testdata/data2.json"
	outFileName = "testdata/data2-out.json"
)

func main() {
	r, err := os.Open(inFileName)	// io.Readerをつくる
	// このファイルには複数のデータが入っている
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	
	var dec *json.Decoder
	dec = json.NewDecoder(r)

	outFile, err := os.Create(outFileName)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(outFile) // bにエンコードした結果を入れる
	// b は bytes.Buffer なので、メモリ上に書く
	
	for dec.More() {
		var o Order
		err := dec.Decode(&o) // 変数oの情報をもとにデコード
		if err != nil {
			log.Fatal(err)
		}
		err = encoder.Encode(o) // tをエンコード ファイルにバイト列で書く
		if err != nil {
			log.Fatal(err)
		}
	}
	outFile.Close()

	// command := fmt.Sprintf("cat \"%s\" ", outFileName)
	result, err := exec.Command("cat", outFileName).Output()
	//	result, err := exec.Command("cat " + outFileName).Output()	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("出力ファイルを表示します\n------------")
	fmt.Printf("%s", result)
	fmt.Println("------------")

	// inFileName と outFileName のファイルが同じになるか確認
	command := "diff -s " + inFileName + " " + outFileName + " | cat"
	result, err = exec.Command("sh", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("diffの結果:\n%s", result)
}
