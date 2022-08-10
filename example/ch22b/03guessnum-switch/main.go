// 数当てゲーム2。少し複雑。関数やループとswitch文を使う

package main  // メインプログラムは必ずこうする

import (  // 読み込むライブラリの宣言
	"fmt"   // 出力関連
	"strconv"  // 文字列から整数などへの変換
	"math/rand"   // 強力なものが必要ならcrypto/randを使う
	"time"  // 乱数のseed生成のため
)

func main() { //liststart1
	answer := getTheNumber();	
loop:  // 下のbreakでループを抜けるための「ラベル」
	for count := 1; ; count++ {  // 無限ループ
		printPrompt(count) // 入力を促すメッセージを表示
		
		switch num, err := readUserAnswer(count); { // 答えを読み込む
		case err != nil || num < 1 || 10 < num:
			fmt.Println("1以上10以下の整数ではないので、ハズレです。")	// fmt.Printlnは改行する
			// breakがなくても、このcaseはここで終了
		case num != answer:
			fmt.Println("残念でした。ハズレです。")
			// breakがなくても、このcaseはここで終了
		default:
			printSuccessMessage(count) // 当たったときのメッセージを表示
			break loop // forループを抜ける。loopは「ラベル」。breakだけだとdefaultを抜けるだけになってしまう
		}
	}
} //listend1

// 当てる数字を得る
func getTheNumber() int {
	rand.Seed(time.Now().UnixNano()) // 乱数のseed。 1970年1月1日午前0時からのナノ秒数
	num := rand.Intn(10)+1  // 0以上10未満の整数をもらって、+1する
	// fmt.Println("答えは: ", num)  // デバッグ用
	return num
}

// プロンプト文字列を表示
func printPrompt(count int) {
	if count == 1 {
		fmt.Print("数当てゲームです。")  // fmt.Printは改行しない
	}
	fmt.Printf("1以上10以下の整数を入力してください（%v回目）: ", count)
	// fmt.PrintfでC言語類似の書式指定可能。%vを指定すると「良きに計らって」くれる
}

// ユーザーからの答えを読み込む
func readUserAnswer(count int) (int, error) {  // 2つの値を返す
	var inp string
	fmt.Scanln(&inp) // 文字列として読み込み   &はポインタを表す（値を書き換えるため）
	return strconv.Atoi(inp) // 整数に変換。errorも返る
}

// 当たったときのメッセージを表示
func printSuccessMessage(count int) { //liststart2
	adverb := ""	
	switch {
	case count == 1:
		fmt.Printf("ビンゴ！ おめでとうございます。一発であたりましたね。素晴らしい！\n")
	case 7 < count:
		adverb = "ヤット"
		fallthrough  // この下のcaseも実行。「落っこちる」
	default:
		fmt.Printf("おめでとうございます。%v回目で%sあたりましたね。\n", count, adverb)
	}
}  //listend2
