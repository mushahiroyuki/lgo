package main

import (
	"context"
	"fmt"
	"time"	
)

func main() {
	ctx := context.Background()
	data := "abc";
	longRunningThingManager(ctx, data)
}

func longRunningThingManager(ctx context.Context, data string) (string, error) { //liststart1
	type wrapper struct {
		result string
		err    error
	}
	ch := make(chan wrapper, 1)
	go func() {
		result, err := longRunningThing(ctx, data)  // 長い処理を実行
		ch <- wrapper{result, err}
	}()
	select {
	case data := <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
} //listend1

func longRunningThing(ctx context.Context, data string) (string, error) {
	fmt.Println("2秒スリープ")
	time.Sleep(2 * time.Second)
	return data, nil
}
