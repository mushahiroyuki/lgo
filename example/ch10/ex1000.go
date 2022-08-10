package main

import "fmt"

func doBusinessLogic(val int) int {
	// valを使って何かをする
	return val*val   // 単純な例：自乗する
}

//liststart1
// ファイルch10/ex1000.go。 詳細は次の節以降で説明
func runThingsConcurrently(chIn <-chan int, chOut chan<- string) { 
	for val := range chIn { // chInから値が到着するたびに
		go func(val int) {  // 新たにゴルーチンを呼び出す
			result := doBusinessLogic(val) // valを処理して結果をresultに代入
			resultString := fmt.Sprintf("%d -> %d", val, result)
			chOut <- resultString // 結果をchOutに入れる
		}(val)
	}
} //listend1

func main() {
	chIn := make(chan int)  // intのチャネルを生成。送信用
	chOut := make(chan string) // stringのチャネルを生成。受信用
	go runThingsConcurrently(chIn, chOut)

	vals := []int{1, 2, 3, 4, 5}
	for v := range vals { // スライスvalsの各要素に対して
		chIn <- v   // その値をchInに送信
	}

	i := 0
	for v := range chOut { // chOutから受信するごとに
		fmt.Println(v) // 結果を出力。順番は不定
		i++
		if len(vals) <= i { // 送った要素の数と同じだけ到着したら
			break  // forループを抜ける
		}
	}

	close(chOut) // chOutをクローズ（終了してしまうので、なくても大丈夫）
	close(chIn) // chOutをクローズ（同様）
}
