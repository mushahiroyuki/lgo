// guessnum2 数当てゲーム2。少し複雑。関数やループとif文を使う
package main  // メインプログラムは必ずこうする

import (  // 読み込むライブラリの宣言
	"fmt"   // 出力関連
	"strconv"  // 文字列から整数などへの変換
	"math/rand"   // 乱数。強力なものが必要ならcrypto/randを使う
	"time"  // 乱数のseed（種）を生成するため
)

// main 
func main() {
	answer := getTheNumber();	// 関数呼び出し。答えの数字をもらう

	for count:=1; ; count++ {  //「条件」を指定しないと無限ループになる
		printPrompt(count) // 説明のメッセージを表示
		
		if num, err := readUserAnswer(count);
		err != nil || num < 1 || 10 < num  {
			// 条件の前で変数を宣言し、値の代入もできる。err以降がifの条件
			fmt.Println("1以上10以下の整数ではないので、ハズレです。")
			// fmt.Printlnは改行する
		} else if num != answer {
			fmt.Println("残念でした。ハズレです。")
		} else {
			printSuccessMessage(count) // あたったときのメッセージを表示
			break  // forループを抜ける
		}
	}
}

// getTheNumber 当てる数字を得る
func getTheNumber() int { // 関数宣言。引数はなし、戻り値はint（整数）
	rand.Seed(time.Now().UnixNano()) //乱数のseed。’70年1月1日0時からのナノ秒数
	num := rand.Intn(10)+1  // 0以上10未満の整数をもらって、+1する
	// fmt.Println("答えは: ", num) 
	return num
}

// printPrompt プロンプト文字列（説明）を表示
func printPrompt(count int) { // 引数は整数。戻り値はなし
	if count == 1 { // 1回目のみ表示
		fmt.Print("数当てゲームです。")  // fmt.Printは改行しない
	}
	fmt.Printf("1以上10以下の整数を（半角で）入力してください（%v回目）: ", count)
	// fmt.Printfではフォーマットが指定できる。%vを指定すると「良きに計らって」くれる
}

// readUserAnswer ユーザーからの答えを読み込む
func readUserAnswer(count int) (int, error) {  // 戻り値が2つある
	var inp string
	fmt.Scanln(&inp) // 文字列として1行読み込み。詳細は前の例のScanlnのところ参照
	return strconv.Atoi(inp) // 整数に変換。errorも戻る
}

// printSuccessMessage 当たったときのメッセージを表示
func printSuccessMessage(count int) {
	if count == 1 {
		fmt.Printf("ビンゴ！ おめでとうございます。一発であたりましたね。素晴らしい！\n")
	} else {
		adverb :="" // 副詞（修飾語「ヤット」をつけるかどうか）
		if  7 < count  {
			adverb = "ヤット"
		}
		fmt.Printf("おめでとうございます。%v回目で%sあたりましたね。\n", count, adverb)
	}
}
