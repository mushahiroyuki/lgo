package main

import (
	"context"
	"errors"
	"fmt"
	"io"   // Go 1.16~
	// "io/ioutil" // Go 1.15まで  1.18でも使えはする
	"net/http"
	"sync"
)

var client = http.Client{} //フィールドの値はすべてゼロ値  //liststart

func callBoth(ctx context.Context, errVal string, slowURL string, fastURL string) {
	ctx, cancel := context.WithCancel(ctx) //キャンセル可能なコンテキストを生成
	defer cancel() // 確実にcancel()呼び出すためにdeferを使う
	
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()

	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()
	
	wg.Wait()
	fmt.Println("両方終了")
}

func callServer(ctx context.Context, label string, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "のリクエスト エラー:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "のレスポンス エラー:", err)
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "のreadエラー:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "の結果:", result)
	}
	if result == "error" {
		fmt.Println("キャンセル:", label)
		return errors.New("エラー発生")
	}
	return nil
} //listend
