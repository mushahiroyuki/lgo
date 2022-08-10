package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"  // 乱数。強力なものが必要ならcrypto/randを使う
)

// この例の目的は、すべてのゴルーチンをきちんと終了すること
// つまりゴルーチンリークをなくすること
func main() {
	rand.Seed(time.Now().UnixNano()) //乱数のseed。’70年1月1日0時からのナノ秒数
	funcs := prepareFunctions() // 関数のスライスを作る
	s := "Time flies like an arrow." // 変換元の文字列
	fmt.Println("オリジナル:", s)	
	r := convertData(s, funcs) // 文字列が戻る
	fmt.Printf("一番早く到着した結果: %v\n\n", r)

	fmt.Println("ほかのゴルーチンが終了するのを待つためにスリープ")	
	time.Sleep(1*time.Second) // 何か処理をしたときの様子を真似するためにsleepで代用。これがないと、ゴルーチンの終了が確認できない
	// この例の目的は、すべてのゴルーチンをきちんと終了すること	
	fmt.Println("main終了")
}


// どの関数が処理しているかわかるように、文字列と関数の番号を一緒にしたものを渡すことにする
type message struct {
	s string
	fromFunc int
}


// 文字列sに関して、convertersに入っている関数を並行に実行して、もっとも早く終了した結果を返す
// 第1引数sは対象の文字列
// 第2引数convertersは「『文字列を受け取って、文字列を返す関数』を要素としてもつスライス」
// 戻り値は文字列
func convertData(s string, converters []func(string) message) message {  
	done := make(chan struct{}) //チャネルの型は空の構造体（3章参照）
	resultChan := make(chan message)
	for _, f := range converters { //convertersに含まれる各関数（f）について
		go func(f func(string) message) { //無名関数をゴルーチンとして起動
			r := f(s)  // 変換を行う
			select {
			case resultChan <- r: // 関数の実行結果をresultChanに送る
				fmt.Printf("結果が戻ってきたのでresultChanに入れたあと: %v\n", r)
			case <-done:
				// 結果が受け取られてしまうとチャネルdoneはクローズされているので
				// doneからゼロ値を受信する（表10.1参照）
				fmt.Println("case <-done選択", r.fromFunc)
			}
			fmt.Println("無名関数終了", r)
		}(f)  // 無名関数の引数。関数をひとつずつ渡す
		
	} // for の終わり
	
	r := <-resultChan // 結果が返ってきたら
	// done <- struct{}{}
	close(done)   // チャネルdoneをクローズ
	return r  
}

func randomPeriod() time.Duration{
	t := time.Millisecond * time.Duration(rand.Intn(4)+2)
	return t
}

func prepareFunctions() []func(string) message  {
	funcNo1 := func(a string) message {
		sleepRandomPeriod(1)
		b := strings.ToLower(a)  // 小文字にする
		fmt.Println("処理完了（1）:", b) // 1: time flies like an arrow.
		return message{b, 1}
	}
	funcNo2 := func(a string) message { // 大文字にする
		sleepRandomPeriod(2)
		b := strings.ToUpper(a) 
		fmt.Println("処理完了（2）:", b) // 2: TIME FLIES LIKE AN ARROW.
		return message{b, 2}
	}
	funcNo3 := func(a string) message { // i -> I
		sleepRandomPeriod(3)
		b := strings.ReplaceAll(a, "i", "I") 
		fmt.Println("処理完了（3）:", b) // 3: TIme flIes lIke an arrow.
		return message{b, 3}
	}
	funcNo4 := func(a string) message { // e->E  r->L
		sleepRandomPeriod(4)
		b := strings.ReplaceAll(a, "e", "E") 
		a = strings.ReplaceAll(a, "r", "L") 
		fmt.Println("処理完了（4）:", b) // 4: 
		return message{b, 4}
	}

	return []func(string) message{funcNo1, funcNo2, funcNo3, funcNo4}
}

func sleepRandomPeriod(funcNum int) {
	t := randomPeriod()
	fmt.Println(funcNum, "will sleep: ", t)
	time.Sleep(t)
}
