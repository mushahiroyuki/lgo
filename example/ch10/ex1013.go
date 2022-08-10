package main

import (
	"fmt"
	"sync"
)

//liststart1
// processAndGather は、スライスdataを受け取って、各要素をfuncで
// 処理した結果を集めて戻す。処理結果の順番はどうなるかわからない
func processAndGather(processor func(int) int, data []int) []int {
	num := len(data)
	chResult := make(chan int, num) // 処理結果を受け取るチャネル
	// dataのサイズと同じ長さのバッファをもつチャネルを作る
	var wg sync.WaitGroup
	wg.Add(num)

	for _, v := range data {
		go func(v int) {
			defer wg.Done() // 処理が終わったらwgをデクリメント
			chResult <- processor(v)  // 処理結果をチャネルに入れる
		}(v)
	}

	wg.Wait()  // すべてのゴルーチンが処理を終了するのを待つ
	close(chResult)

	var result []int
	for v := range chResult {  // chResultのバッファに残っている要素を処理
		result = append(result, v)
	}
	return result
}
//listend1

func main() {
	inpValues := []int{1, 2, 3, 4, 5}
	outValues := processAndGather(
		func(j int) int {return j*j}, // 第1引数は関数
		inpValues)  // 第2引数は計算する値（intのスライス）
	
	fmt.Println(outValues)
}
