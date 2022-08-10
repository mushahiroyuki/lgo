package main

import(
	"fmt"
	"net/http" // https://pkg.go.dev/net参照。これを見ずに理解は不可能か
	"time"
	"context"
	"encoding/json"
)

func main() {
	client := http.Client{ //liststart1
		Timeout: 30 * time.Second, // タイムリミットの設定  30秒
	}  //listend1

//liststart2
	url := "https://jsonplaceholder.typicode.com/todos/1" //JSONデータを返してくれる無料サイト
	req, err := http.NewRequestWithContext(
		context.Background(), // 12章参照。当面はこのまま利用
		http.MethodGet, // HTTPのメソッド"GET"に対応する定数
		url,
		nil) // io.Readerでボディを指定（この場合はなし）
	if err != nil {
		panic(err)
	} //listend2

	req.Header.Add("X-My-Client", "Learning Go") //liststart3
	// リクエストのヘッダに追加。この場合はなくても結果は同じ...
	res, err := client.Do(req) // 準備したリクエストを送信し結果をもらう
	if err != nil {
		panic(err)
	} //listend3

	
	defer res.Body.Close() //liststart4
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data) // %+vでフィールド名付き表で示
//listend4
}
