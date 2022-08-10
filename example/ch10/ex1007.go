package main

import (
	"fmt"
	"strings"
	"time"
)

// この例の目的は、すべてのゴルーチンをきちんと終了すること
// つまりゴルーチンリークをなくすること
// ex1007b.goにより細かく説明を出力するバージョンがあります
// 原著と合わせる都合上searcherとなっていますが、検索するわけではなく
// converter といった名前のほうがよいと思いますがあしからずご了承ください

func main() {
	funcs := prepareFunctions() // 関数のスライスを作る
	
	s := "Time flies like an arrow." // 変換元の文字列
	r := searchData(s, funcs) // スペースで区切った文字列のスライスが戻る
	fmt.Println("結果:", r)

	time.Sleep(1*time.Second) // 何か処理をしたときの様子を真似するためにsleepで代用。これがないと、ゴルーチンの終了が確認できない（コメントアウトして実行してみてください。上のimportのtimeの前に "_ "も必要です）
	fmt.Println("mainを終了")
}


//liststart1
// 文字列sに関して、searchersに入っている関数を並行に実行して、もっとも早く終了した結果を返す
// 第1引数sは対象の文字列
// 第2引数searchersは「『文字列を受け取って、文字列のスライスを返す関数』を要素としてもつスライス」
// 戻り値は文字列のスライス
func searchData(s string, searchers []func(string) []string) []string {  
	done := make(chan struct{}) // チャネルの型は空の構造体（3章参照）
	resultChan := make(chan []string)
	for _, searcher := range searchers {
		// 以下をsearchersに含まれる各関数（searcher）について行う

		// 無名関数をゴルーチンとして起動。引数は「文字列を受け取って、
		// 文字列のスライスを返す関数」
		go func(f func(string) []string) { 
			select {
			case resultChan <- f(s): // 関数の実行結果をresultChanに送る
				fmt.Println("結果が戻ってきた")
			case <-done:
				fmt.Println("doneを選択")
			}
		}(searcher)  // 無名関数の引数。関数をひとつずつ渡す
		
	} // for の終わり
	
	r := <-resultChan // 結果が返ってきたら
	close(done)   // チャネルdoneをクローズ
	return r  // 文字列のスライス
} //listend1


func prepareFunctions() []func(string) []string {
	searcher1 := func(a string) []string {
		b := strings.ToLower(a)  // 小文字にして
		fmt.Println("1:", b) // 1: time flies like an arrow.
		r := strings.Split(b, " ") // スペースで分割
		return r
	}
	searcher2 := func(a string) []string {
		b := strings.ToUpper(a) 
		fmt.Println("2:", b) // 2: TIME FLIES LIKE AN ARROW.
		r := strings.Split(b, " ") // スペースで分割
		return r
	}

	searcher3 := func(a string) []string {
		b := strings.ReplaceAll(a, "i", "I") 
		fmt.Println("3:", b) // 3: TIme flIes lIke an arrow.
		r := strings.Split(b, " ") // スペースで分割
		return r
	}

	searcher4 := func(a string) []string {
		b := strings.ReplaceAll(a, "e", "E") 
		a = strings.ReplaceAll(a, "r", "L") 
		fmt.Println("4:", b) // 4: 
		r := strings.Split(b, " ") // スペースで分割
		return r
	}
	// 関数のスライスを作る
	funcs := []func(string) []string{searcher1, searcher2, searcher3, searcher4}
	return funcs
}
