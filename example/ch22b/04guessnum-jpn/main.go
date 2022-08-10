// 数当てゲーム3。日本語の識別子を使う

package main  // メインプログラムは必ずこうする

import (  // 読み込むライブラリの宣言
	"fmt"   // 出力関連
	"strconv"  // 文字列から整数などへの変換
	"math/rand"   // 強力なものが必要ならcrypto/randを使う
	"time"  // 乱数のseed生成のため
)

func main() { //liststart1
	正解 := 当てる数字を決定();	
loop:
	for カウンタ := 1; ; カウンタ++ {  // 無限ループ
		プロンプトを表示(カウンタ) // 入力を促すメッセージを表示

		switch 解答, err := ユーザーからの答えを取得(カウンタ); { // 答えを読み込む
		case err != nil || 解答 < 1 || 10 < 解答:
			fmt.Println("1以上10以下の整数ではないので、ハズレです。")
			// breakがなくても、このcaseはここで終了
		case 解答 != 正解:
			fmt.Println("残念でした。ハズレです。")
			// breakがなくても、このcaseはここで終了
		default:
			成功時のメッセージを表示(カウンタ) // 当たったときのメッセージを表示
			break loop // forループを抜ける。breakだけだとdefaultを抜けるだけになる
		}
	}
} //listend1

func 当てる数字を決定() int {
	rand.Seed(time.Now().UnixNano()) // 乱数のseed。 1970年1月1日午前0時からのナノ秒数
	n := rand.Intn(10)+1  // 0以上10未満の整数をもらって、+1する
	// fmt.Println("答えは: ", n)  // デバッグ用
	return n
}

func プロンプトを表示(カウンタ int) {
	if カウンタ == 1 {
		fmt.Print("数当てゲームです。")  // fmt.Printは改行しない
	}
	fmt.Printf("1以上10以下の整数を入力してください（%v回目）: ", カウンタ)
	// fmt.PrintfでC言語類似の書式指定可能。%vを指定すると「良きに計らって」くれる
}

func ユーザーからの答えを取得(カウンタ int) (int, error) {  // 2つの値を返す
	var 入力 string
	fmt.Scanln(&入力) // 文字列として読み込み   &はポインタを表す（値を書き換えるため）
	return strconv.Atoi(入力) // 整数に変換。errorも返る
}

func 成功時のメッセージを表示(カウンタ int) { //liststart2
	修飾語 := ""	
	switch {
	case カウンタ == 1:
		fmt.Printf("ビンゴ！ おめでとうございます。一発であたりましたね。素晴らしい！\n")
	case 7 < カウンタ:
		修飾語 = "ヤット"
		fallthrough  // この下のcaseも実行。「落っこちる」
	default:
		fmt.Printf("おめでとうございます。%v回目で%sあたりましたね。\n", カウンタ, 修飾語)
	}
}  //listend2
