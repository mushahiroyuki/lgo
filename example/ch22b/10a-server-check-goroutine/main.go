//
package main

import(
	"fmt"
	"net/http"
	"time"
)

const TimeLimit time.Duration = 200 * time.Millisecond
// タイムリミットを設定
// const TimeLimit time.Duration = 80 * time.Millisecond

type ServerStatus int // サーバの状況を示す型ServerStatusを定義

const ( // ServerStatusが取りうる具体的な値（定数）の定義
	不明   ServerStatus = iota  // 「不明」に 整数の0が割り当てられる
	アップ    // 「アップ」に1
	問題発生  // 「問題発生」に2
)
// ↑変数に使える文字なら日本語も使える

type site struct {
	url string // URL文字列
	status ServerStatus  // 現状。「不明」「アップ」「問題発生」のいずれか
}


func main() {
	websites := []site {  // URLと現状。構造体のスライス
		{"https://www.oreilly.co.jp/", 不明},
		{"https://musha.com/", 不明}, 
		{"https://marlin-arms.com/", 不明},
		{"https://dictjuggler.net/", 不明},
		{"http://localhost/", 不明},
	}
	
	ch := make(chan site) // siteをやり取りするチャネルchを作る
	defer close(ch) // 最後にチャネルを閉じる
	client := &http.Client{ // 「11.4.1 クライアント」参照
		Timeout: time.Duration(TimeLimit), // タイムリミットの設定
	}
	for _, site := range websites { // 各サイトについて
		go checkWebsite(site, ch, client) // ゴルーチンを生成
	}

	for i:=0; i<len(websites); i++ { // サイトの数だけループ
		siteResponded := <-ch // どのサイトから返事が来るか、順番はわからない
		fmt.Printf("%s: %v\n", siteResponded.url, siteResponded.status)
		// チャネル経由で返事が来たサイトのURLと状況を書く
	}
}

// checkWebsite siteに指定されたページをチェックする
func checkWebsite(s site, ch chan<-site, client *http.Client) {	
	if _, err := client.Get(s.url); err != nil { // urlから
		s.status = 問題発生
	} else {
		s.status = アップ
	}
	ch <- s
}

// ServerStatus 専用の関数（インタフェース）String。これを定義しておくことで、
// fmt.Printfに%vを指定した時に、数字ではなくここで指定した文字列で表示される
// 実行結果参照
func (ss ServerStatus) String() string {
	switch ss {
	case 不明:
		return "不明"
	case アップ:
		return "OK"
	case 問題発生:
		return "★問題発生★"
	default: // これがあれば、将来状態を追加しても何かが表示される
		return fmt.Sprintf("%d", int(ss))
	}
}
