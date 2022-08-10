package main

import(
	"fmt"
	"sync"
	"net/http"
	"time"
)

func main() {
	var wg sync.WaitGroup // WaitGroupを使って終了してもよい時を判断
	sites := []string{  // チェックするサイトのURL。文字列のスライス
		"https://www.oreilly.co.jp/",
		"https://musha.com/",
		"https://marlin-arms.com/",
		"https://dictjuggler.net/",
		"http://localhost/",

		"https://www.xxxxxxyyyyyzzzz.jp/",	// 存在しないサイト
		"https://musha.com/non", // 存在しないページ
		"https://musha.com/03support/learning-go/response/moved-temp/",//redirect
		"https://musha.com/03support/learning-go/response/moved-perm/",//redirect
	}

	client := &http.Client{
    Timeout: 200 * time.Millisecond, // タイムアウトの設定  ミリ秒単位 11章など参照
		CheckRedirect: nil,
	}

	for i, s := range sites { // 各サイトについて
		go checkWebsite(s, i, &wg, client) // ゴルーチンを生成
		wg.Add(1) // WaitGroupのカウンタを1増やす（処理が終了したら1減る）
	}

	wg.Wait() // WaitGroupのカウンタが0になるのを待つ
	// 0になれば全部確認が終わっているので、終了する
}

// checkWebsite 指定されたページを確認
func checkWebsite(url string, i int, wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done() // 抜ける時にWaitGroupのカウンタを1減らす
	// 忘れないように冒頭にdeferを使って書いておく
	if res, err := client.Get(url); err != nil {
		fmt.Printf("%d %s  %v\n", i, url, err)
	} else {
		fmt.Printf("%d %s  code: %v\n", i, url, res.Status)
	}
}
