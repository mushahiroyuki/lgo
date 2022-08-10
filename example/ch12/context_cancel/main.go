package main

import (
	"context"
	"os"
)

func main() { //liststart
	ss := slowServer() // サーバを起動
	defer ss.Close()
	fs := fastServer()
	defer fs.Close()

	ctx := context.Background() // コンテキストを生成
	callBoth(ctx, os.Args[1], ss.URL, fs.URL)
}  //listend

