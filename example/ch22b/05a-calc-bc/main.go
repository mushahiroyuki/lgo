// 簡易電卓 -- コマンドライン引数、正規表現、文字列置換の例
package main

import (
	"fmt"
	"os" // コマンドライン引数の処理
	"strings"  // 文字列置換
	"regexp"  // 正規表現
	"bufio" // 読み込み
	"golang.org/x/text/width" // 全角 -> 半角変換
)

func main() {
	var exp string  // EXPression: 式
	if n := len(os.Args); 2 <= n { // 引数があったら
		for i := 1; i < n; i++ {
			exp += os.Args[i] // 引数を後ろに追加していく
		}
		calculateAndPrintValue(exp) // 必要な変換を行って計算する
	}

	scanner := bufio.NewScanner(os.Stdin) // 標準入力を受け付ける「スキャナ」
	for scanner.Scan() { // １行分の入力を取得できる限り繰り返す
		exp = scanner.Text()  // 入力のテキスト（string）を取得
		if exp == "" {
			return  // 空行入力で終了
		}
		calculateAndPrintValue(exp) // 必要な変換を行って計算する
	}
}

// calculateAndPrintValue  文字の置換を行って計算し出力する
func calculateAndPrintValue(exp string) error {
	expOrigin := exp
	exp = replaceChars(exp) // ÷->/  全角->半角変換などを行う
	if (exp != expOrigin) {
		fmt.Println("次の計算をします:", exp) // 変えた時は確認のため表示
	}
	result, err := calculate(exp) // 計算実行。エラーが起こればerrがnilでなくなる
	if err != nil || len(result)==0 { // エラーが起こったか、結果が空の時
		fmt.Println("計算できません")
	} else {
		fmt.Printf("%s\n", result)  
	}
	return err
}

// replaceChars  置換する
func replaceChars(exp string) string {
	exp = width.Fold.String(exp)  // 全角 -> 半角
	re := regexp.MustCompile(`[xX]`)
	exp = re.ReplaceAllString(exp, "*") // 「x」「X」 -> 「*」

	re = regexp.MustCompile(`[,、，]`) // 3桁ごとの区切り
	exp = re.ReplaceAllString(exp, "") // 「,」「、」「，」 削除

	re = regexp.MustCompile(`[。．]`)
	exp = re.ReplaceAllString(exp, ".") //全角の 「．」「。」 -> 半角の「.」

	exp = strings.ReplaceAll(exp, "÷", "/") //  「÷」 -> 「/」
	return exp
}
