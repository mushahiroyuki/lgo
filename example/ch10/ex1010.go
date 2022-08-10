/*
ex1010.sh を実行すると12個のリクエストを送ります。
*/

package main

import (
 	"fmt"
	"errors"
	"time"
	"net/http"
)

// pressure gaugeは「圧力計」のこと
type PressureGauge struct { //liststart1
	ch chan struct{} // 空の構造体
}

// サイズlimitのpressure gaugeを戻す
func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}  // limit個のインスタンス（空の構造体）を入れる
	}
	return &PressureGauge{
		ch: ch, // 満杯になっているチャネルを戻す
	}
}


// PressureGauge pgを受け取って、それが許す範囲で関数fを実行する
func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()  // fを実行する
		pg.ch <- struct{}{} // ひとつ余裕ができた。空の構造体のインスタンスを送る
		return nil
	default:
		return errors.New("キャパシティに余裕がありません")
	}
} //listend1

func doThingThatShouldBeLimited() string { //liststart2
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := New(10) // PressureGaugeの初期化。最大数を10個にする
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() { // 無名関数をpgの範囲で処理
			w.Write([]byte(doThingThatShouldBeLimited())) // バイトのスライスに型変換
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	fmt.Println("ブラウザで次を開いてください: 'http://localhost:8080/request'")
	fmt.Println("あるいは 'sh ex1010.sh' を実行してみてください")	
	http.ListenAndServe(":8080", nil)
} //listend2
