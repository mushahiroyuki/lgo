package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 { // ファイル名が指定されている？
		log.Fatal("ファイルが指定されていません")
	}
	f, err := os.Open(os.Args[1]) // ファイルをオープン
	if err != nil {
		log.Fatal(err) // オープンに問題あり。エラーを出力して終了
	}
	defer f.Close() // 後始末のコード

	data := make([]byte, 2048) // バイトのスライスを生成
	for { // 無限ループ
		count, err := f.Read(data) // 読み込んだバイト数とエラーを返す
		os.Stdout.Write(data[:count]) // 「標準出力」に出力
		if err != nil {
			if err != io.EOF { // ファイルの終わりでないならば
				log.Fatal(err)  // エラーを出力して終了
			}
			break  // forループを抜ける
		}
	}
}
