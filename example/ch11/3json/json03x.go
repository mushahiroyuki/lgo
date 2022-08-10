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
	// 一時ファイルの作成。JSON形式でここに書き込む
	// func CreateTemp(dir, pattern string) (*File, error)
	//   tmpFile の型は *File。*Fileはio.Writer（およびio.Reader）を満たす
	
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name()) // 使い終わったら削除
	encoder := json.NewEncoder(tmpFile) // *Encoderが返る
	// tmpFileにエンコードする
	// fmt.Printf("encoderの型: %T\n", encoder) // *json.Encoder
	// func NewEncoder(w io.Writer)  // encoding/jsonパッケージ
	err = encoder.Encode(dataToFile)	  //  dataToFileのデータをエンコード
	//  func (enc *Encoder) Encode(v any) error
	// encoderは Encoder型なので、メソッドEncodeを起動できる
	// EncodeはdataToFileに入っているデータをJSONにコード化したものを
	// ストリーム（tmpFile）に書き出し最後に改行文字を付加する
	// encoderがEncodeすると、ストリームに書き出される点に注意！
	// 単にコード化するだけではない

	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	} //listend2

	// 読み込み
	tmpFile2, err := os.Open(tmpFile.Name()) //liststart3
	if err != nil {
		panic(err)
	}
	var dataFromFile Person
	// err = json.NewDecoder(tmpFile2).Decode(&dataFromFile)

	var decoder *json.Decoder
	decoder = json.NewDecoder(tmpFile2)
	err = decoder.Decode(&dataFromFile) // detaFromFileにデコード
	
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("dataFromFile: %+v\n", dataFromFile) // 構造体をフィールド名付きで表示
	// JSON形式の一時ファイルから構造体に読み込んだ内容を表示
//listend3	
}
