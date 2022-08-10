package main

import (
	"fmt"
	"strings"
	"time"
	"math/rand"  // 乱数。強力なものが必要ならcrypto/randを使う
)

type strConv func(string) string // 文字列変換の関数

func main() {
	rand.Seed(time.Now().UnixNano()) //乱数のseed。’70年1月1日0時からのナノ秒数

	strConvFuncs := []strConv{strings.ToLower, strings.ToUpper}
	s := "Time flies like an arrow." // 変換元の文字列
	fmt.Println("オリジナル:", s)

	// 	文字列 s を strConvFuncs で変換した結果を集める
	//  オリジナルのsも含め、すべてをスライスrに入れる
	
	r := []string{s}  // 最終結果。最初の要素は s
	r2 := getVariations(s, strConvFuncs)
	r = append(r, r2...) // スライスを後ろに追加
	fmt.Println("結果:", r)
}


func getVariations(s string, converters []strConv) []string {
	done := make(chan struct{}) //doneチャネル。doneチャネルの例参照
	chs := []chan string{make(chan string), make(chan string)}

	for i, f := range converters {
		go func(s string, f strConv, ch chan string, i int) {
			s = f(s)
			fmt.Printf("変換結果 %d: %s\n", i, s)
			ch <- s
			close(ch)
		}(s, f, chs[i], i)
	}

	r := []string{} // 結果のスライス。変化形をもどす
	// select は クローズされているチャネルを読むといつも成功しゼロ値が戻る
	for {
		select {
		case v, ok := <- chs[0]: // 1回目はvに入る。2回目はcloseされている
			if !ok { // inはクローズされていた
				chs[0] = nil // nilチャネルからは何も読み込まない。したがってこのcaseは再度成功することはない
				fmt.Println("continueする: chan0")
				continue
			}
			fmt.Println("変換結果をappend:", v)
			r = append(r, v)
		case v, ok := <- chs[1]:
			if !ok { // in2はクローズされていた
				chs[1] = nil // このcaseは再度成功することはない
				fmt.Println("continueする: chan1")
				continue
			}
			fmt.Println("appending:", v)
			r = append(r, v)
		case <-done:
			fmt.Println("case <-done選択")
			return r
		}

		if len(r)>=len(chs) { // 全部答えが来た
			close(done)
		}

	}
}

func randomPeriod() time.Duration{
	t := time.Millisecond * time.Duration(rand.Intn(4)+2)
	return t
}


func sleepRandomPeriod(funcNum int) {
	t := randomPeriod()
	fmt.Println(funcNum, "will sleep: ", t)
	time.Sleep(t)
}
