// 簡易電卓。コマンドライン引数、正規表現、文字列置換

package main

import (
	"fmt"
  "go/token" // Evalに必要
  "go/types" // 関数Eval（文字列を評価: EVALuateする）のため
_	"os/exec" // exec.Command
_  "math"     // 冪上の計算
	"os"      // コマンドライン引数の処理
	"regexp"  // 正規表現
	"strings" // 文字列変換
	"golang.org/x/text/width" // 全角 -> 半角変換
)

func main() {
	var exp string // EXPression: 式
	if n := len(os.Args); 2 <= n {
		for i := 1; i < n; i++ {
			exp += os.Args[i] // 引数を後ろに追加していく
		}
		calculateAndPrintValue(exp)
	}

	for {
		exp = readExpression()
		if exp == "" {
			return // 空行入力で終了
		}
		calculateAndPrintValue(exp)
	}

}

func readExpression() string {
	var inp string
	fmt.Scanln(&inp) // 文字列として読み込み   &はポインタを表す（値を書き換えるため）
	fmt.Println(inp)
	return inp
}

func calculateAndPrintValue(exp string) error {
	exp = replaceChars(exp) // ÷->/  全角->半角変換などを行う
	fmt.Println("次の計算をします:", exp)

	tv, err := types.Eval(token.NewFileSet(), nil, token.NoPos, exp)
	if err != nil {
		fmt.Println("計算できません ── ")
	}
	fmt.Println(tv.Value)
	return err
}

func replaceChars(exp string) string {
	exp = width.Fold.String(exp)           // 全角 -> 半角
	exp = strings.ReplaceAll(exp, ",", "") // 3桁ごとの区切りの「,」を削除

	//	fmt.Println("$1: ", re1)
	//	fmt.Println("$2: ", "$2")
	//	val := math.Pow(, "$2")
	//	exp = re1.ReplaceAllString(exp, val)
	// fmt.Println("途中経過:", exp)
	// math.Pow(べき上を計算する数値,べき乗)

	re := regexp.MustCompile(`[xX]`)
	exp = re.ReplaceAllString(exp, "*")

	exp = strings.ReplaceAll(exp, "÷", "/")

	return exp
}
