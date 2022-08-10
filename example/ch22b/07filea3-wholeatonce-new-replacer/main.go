// すべてのデータを一度に読み込んで処理。bytes.ReplaceAllを使う
package main

import (
	"fmt"
	"os" // ファイルの読み込みなど
	"strings"
	"log"  // ログを書くのに便利。
)

// 定数の宣言（詳細は「2.3 定数」）。型は右辺のデフォルト（この場合string）になる
const path1 = "testdata/dict.txt"

func main(){
	allData, err := os.ReadFile(path1) // 全データを一度に読み込み
	if err != nil {
		log.Fatal(err) // エラーをログに書いて異常終了
	}
//liststart	
	r := strings.NewReplacer("//", "##", "\t", "|") // 一度に複数のペアを指定できる
	result := r.Replace(string(allData)) // replacerを使って文字列を置換
	fmt.Printf("%s", result)  //listend
}

