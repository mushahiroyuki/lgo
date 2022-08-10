package main

import (
	"fmt"
	"time"
	"math/rand"
)

func processChannel(ch  chan int) []int { //liststart1
	const maxConc = 10
	results := make(chan int, maxConc)
	for i := 0; i < maxConc; i++ {
		go func() {
			v := <- ch
			results <- process(v)
		}()
	}
	fmt.Println("ゴルーチン 起動完了")
	
	var out []int // intのスライス
	for i:=0; i<maxConc; i++ {
		out = append(out, <-results) // 結果を受け取って後ろに追加
	}
	return out
}
//listend1

func process(v int) int {	
	returnVal := v*v
	rand.Seed(time.Now().UnixNano()) // シードの指定
	sleepSec := rand.Intn(3) // 0以上3未満の整数を戻す

	fmt.Println("process:", v, returnVal, sleepSec)
	time.Sleep(time.Duration(sleepSec)*time.Second)
	return returnVal
}


func main() {
	ch := make(chan int)

	var result []int;

	go func() {  // 処理してもらう数値をchに入れる
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	
	result = processChannel(ch)
	
	fmt.Printf("result: %d\n", result)
}
