// すべてのデータを一度に読み込んで処理。bytes.ReplaceAllを使う
package main

import (
	"fmt"
	"os" // ファイルの読み込みなど
	"bytes" // バイト列のまま文字列置換
	"log"  // ログを書くのに便利。
)

// 定数の宣言（詳細は「2.3 定数」）。型は右辺のデフォルト（この場合string）になる
const path1 = "testdata/dict.txt"

func main(){
	allData, err := os.ReadFile(path1) // 全データを一度に読み込み
	// allDataは[]uint8（8ビット符号なし整数のスライス）
	if err != nil {
		log.Fatal(err) // エラーをログに書いて異常終了
	}
	allData = bytes.ReplaceAll(allData, []byte( "//"), []byte("##")) //liststart
	// 第2、第3引数はバイト列に変換してから呼び出し
	allData = bytes.ReplaceAll(allData, []byte( "\t"), []byte("|")) 
	fmt.Printf("%s", allData) // 置換した結果を出力。%sを指定すれば文字列になる //listend
	// fmt.Print(allData) // ←これだと、バイト列として出力 
}
