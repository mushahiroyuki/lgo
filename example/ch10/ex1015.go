package main

import (
	"fmt"
	"context"
	"math/rand"  // 乱数。強力なものが必要ならcrypto/randを使う
	"time"
	"os"
	"log"
)

type AOut int
type BOut int
type COut int
type CIn struct {
	A AOut
	B BOut
}
type Input int

type processor struct { // 全体の処理の構成 //liststart1
	outA chan AOut
	outB chan BOut
	outC chan COut
	inC  chan CIn
	errs chan error
} //listend1

const maxInt = 20 //liststart2
const timeLimit = 50

// AとBからデータを集めて（gather）、Cに処理（process）してもらう
func gatherAndProcess(ctx context.Context, data Input) (COut, error) {
////	func processAndGather(ctx context.Context, data Input) (COut, error) {
////	func GatherAndProcess(ctx context.Context, data Input) (COut, error) {
//// ch := make(chan int)
	log.SetFlags(log.Lmicroseconds) //logのタイムスタンプをマイクロ秒単位にする
	log.Println("staring; timeLimit:", timeLimit*time.Millisecond)
	ctx, cancel := context.WithTimeout(ctx, timeLimit*time.Millisecond) //G1
	// コンテキストのセットアップ。Doneチャネルとキャンセレーション関数を戻す
	defer cancel() //G2
	p := processor{ //G3
		outA: make(chan AOut, 1),
		outB: make(chan BOut, 1),
		inC:  make(chan CIn, 1),
		outC: make(chan COut, 1),
		errs: make(chan error, 2), //G4
	}
	p.launch(ctx, data) //G5
	inputC, err := p.waitForAB(ctx) //ABからの出力を待つ //G6  
	if err != nil { // エラーが起こったら
		fmt.Println("Error returned from waitForAB:", err)
		os.Exit(1) // プログラムを終了
	}
	p.inC <- inputC //G7
	out, err := p.waitForC(ctx) //G8
	return out, err //G9
} //listend2

func (p *processor) launch(ctx context.Context, data Input) { //liststart3
	go func() { //L1
////		aOut, err := getResultA(ctx, data.A)
		aOut, err := getResultA(ctx, data)
		if err != nil {
			p.errs <- err //L2
			return
		}
		p.outA <- aOut //L3
	}()

	go func() { //L4
////		bOut, err := getResultB(ctx, data.B)
		bOut, err := getResultB(ctx, data)
		if err != nil { //L5
			p.errs <- err
			return
		}
		p.outB <- bOut //L6
	}()

	go func() {
		select { //L7
		case <-ctx.Done(): //L8
			log.Println("case ctx.Done()")
			return
		case inputC := <-p.inC: //L9
			cOut, err := getResultC(ctx, inputC) //L10
			if err != nil {
				p.errs <- err
				return
			}
			p.outC <- cOut
		}
	}()
} //listend3

func (p *processor) waitForAB(ctx context.Context) (CIn, error) { //liststart4
	var inputC CIn
	count := 0
	for count < 2 {
		select {
		case a := <-p.outA: //W1
			inputC.A = a //W2
			count++
		case b := <-p.outB: //W3
			inputC.B = b //W4
			count++
		case err := <-p.errs: //W5
			return CIn{}, err
		case <-ctx.Done(): //W6
			return CIn{}, ctx.Err()
		}
	}
	return inputC, nil //W7
} //listend4

func (p *processor) waitForC(ctx context.Context) (COut, error) { //liststart5
	select { //C1
	case out := <-p.outC: //C2
		return out, nil
	case err := <-p.errs: //C3
		return 0, err
		//// COut{}
	case <-ctx.Done(): //C4
		return 0, ctx.Err()
		////COut{}
	}
} //listend5

func main() {
	rand.Seed(time.Now().UnixNano()) //乱数のseed。’70年1月1日0時からのナノ秒数
	ctx := context.Background()

	n := Input(rand.Intn(maxInt)+1)
	v, err := gatherAndProcess(ctx, n)
	//// 	values := processAndGather(ctx, sites)
	if (err != nil) {
		fmt.Println("Error: ", err)
		return
	}
	log.Println("Final Result:", v)
}

func getResultA(ctx context.Context, data Input) (AOut, error) {
	time.Sleep(randomPeriod("A", 1))
	r := AOut(data * Input(rand.Intn(maxInt)+1)) // 型変換
	log.Println("from A:", r)
	return r, nil
}


func getResultB(ctx context.Context, data Input) (BOut, error) {
	time.Sleep(randomPeriod("B", 1))
	r := BOut(data * Input(rand.Intn(maxInt)+1)) // 型変換
	log.Println("from B:", r)	
	return r, nil
}

func getResultC(ctx context.Context, data CIn) (COut, error) {
	time.Sleep(randomPeriod("C", 3))
	r := int(data.A) + int(data.B)
	return COut(r), nil
}


func randomPeriod(which string, div int) time.Duration {
	i := rand.Intn(timeLimit) / div + 1 + timeLimit/10
	t := time.Millisecond * time.Duration(i)
	log.Println(which, "will sleep:", t)
	return t
}

