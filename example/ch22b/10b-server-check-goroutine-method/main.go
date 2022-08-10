// ゴルーチン版サーバチェッカー。メソッドを定義
package main

import(
	"fmt"
	"net/http"
	"time"
)

const TimeLimit time.Duration = 200 * time.Millisecond // 制限時間の設定
// const TimeLimit time.Duration = 80 * time.Millisecond

type ServerStatus int // サーバの状況を示す型ServerStatusを定義

const ( // ServerStatusが取りうる具体的な値（定数）の定義
	不明   ServerStatus = iota  // 「不明」に 整数の0　が割り当てられる
	アップ    // 「アップ」に1
	問題発生  // 「問題発生」に2
) // 識別子には日本語も使えます

// ServerStatus 専用の関数（インタフェース）String。これを定義しておくことで、
// fmt.Printfに%s（%v）を指定した時に、数字ではなくここで指定した文字列で表示される
// 詳しくは「7.2 メソッド」参照
func (ss ServerStatus) String() string {
	switch ss {
	case 不明:
		return "不明"
	case アップ:
		return "OK"
	case 問題発生:
		return "★問題発生★"
	default: // ここには来ないはずだが念の為
		return fmt.Sprintf("%d", int(ss))
	}
}

// Site ウェブサイトを表す構造体
type Site struct {  
	url string // URL文字列
	status ServerStatus  // サーバの現状。「不明」「アップ」「問題発生」のいずれか
}

// upDownChecker 型Siteに付随するメソッド（型SiteにメソッドupDownCheckerを付加）
// Siteの状態をチェック
func (s Site) upDownChecker(ch chan<-Site) {
	client := &http.Client{ // 「11.4.1 クライアント」参照
		Timeout: time.Duration(TimeLimit), // タイムリミットの設定
	}
	
	if _, err := client.Get(s.url); err != nil {
		s.status = 問題発生
	} else {
		s.status = アップ
	}
	ch <- s  // ゴルーチンとして起動されるのでチャネルを経由して返す
}

func main() {
	websites := []Site {  // URLと現状
		{"https://www.oreilly.co.jp/", 不明},
		{"https://musha.com/", 不明}, 
		{"https://marlin-arms.com/", 不明},
		{"https://dictjuggler.net/", 不明},
		{"http://localhost/", 不明},
	}
	
	ch := make(chan Site) // Siteをやり取りするチャネルchを作る
	defer close(ch) // 最後にチャネルを閉じる
	for _, site := range websites { // 各サイトについて
		go site.upDownChecker(ch) // ゴルーチンを生成。結果はチャネルch経由で返してもらう
	}

	for i:=0; i<len(websites); i++ { // サイトの数だけループ
		siteResponded := <-ch // どのサイトから返事が来るか、順番はわからない
		fmt.Printf("%s: %s\n", siteResponded.url, siteResponded.status)
		// チャネル経由で返事が来たサイトのURLと状況を書く
	}
}


