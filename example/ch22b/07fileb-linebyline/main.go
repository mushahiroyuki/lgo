// 1行ずつ正規表現を使って処理
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
)

var re1 = regexp.MustCompile(`[,，、] *`) // 正規表現をコンパイルしておく

func main(){
	data, _ := os.Open("testdata/dict2.txt")
	defer data.Close()
	
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" || s[0:1] == "#" { // s[0:1] で先頭の1文字
			continue  // この行はスキップ（コメントあるいは空白行）
		}
		columns := strings.Split(s, "|") // "|"で分割
		if len(columns)!=2 { // カラム数が2個でないとき
			continue // この行はスキップ。形式がおかしい
		}
		eng := columns[0] // 英語部分
		jpn := columns[1] // 日本語訳
		translations := re1.Split(jpn, -1) // 正規表現を使って分割。-1は全部
		// []string（文字列のスライス）が戻る
		for _, translation := range translations {
			// range からはインデックス（添字）と文字列が戻るが、インデックスは無視する
			// 「_」に代入すると無視できる
			// 変数に代入するとその変数を使わないとエラーになってしまうので
			fmt.Printf("%s|%s\n", eng, translation);
		}
	}
}
