package main

import (
	"fmt"
	"os/exec" // OSコマンドの実行
	"strings"
)
// calculate  bcコマンドを呼び出して計算する
func calculate(exp string) (string, error) {
	command := fmt.Sprintf("echo \"%s\" | bc -l", exp)
	// コマンドの文字列を作る。Sprintfの%sの部分がexpの値で置き換わる
	result, err := exec.Command("sh", "-c", command).Output()
	// resultはバイト列（型byteのスライス -- []byte）
	// スライスは「可変長の配列」。Goの配列はサイズ固定だがスライスは変更可
	// err はエラーがあるかどうかを示す。errが nilならば正常終了
	// err の処理は呼び出し側（calculateAndPrintValue）が行う
	if err != nil {
		return "", err
	}
	r := string(result)
	r = strings.TrimRight(r, "\n") // 最後に改行がついて戻るので削除する
	return r, err // 2つの値を戻す
}
