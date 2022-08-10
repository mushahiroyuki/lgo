// すべてのデータを一度に読み込んで処理
package main

import (
	"fmt"
	"os" // ファイルの読み込みなど
	"strings" // 文字列置換
	"log"  // ログを書くのに便利。
)

// 定数の宣言（詳細は「2.3 定数」）。型は右辺のデフォルト（この場合string）になる
const path1 = "testdata/dict.txt"

func main(){
	allData, err := os.ReadFile(path1) // 全データを一度に読み込み	
	// allDataは[]uint8（8ビット符号なし整数のスライス）
	// io/ioutilにもReadFileがある。現在も動作するが、今は単に上のos.ReadFileが呼ばれる
	if err != nil {
		log.Fatal(err) // エラーをログに書いて異常終了
	}
	s := string(allData) 	// string(allData)で型をstringに変換してから呼び出し
	s = strings.ReplaceAll(s, "//", "##") // 「//」を「##」にすべて置換
	s = strings.ReplaceAll(s, "\t", "|")  // 「タブ」を「|」にすべて置換
	fmt.Printf("%s", s) // 文字列置換した結果を出力	
}
