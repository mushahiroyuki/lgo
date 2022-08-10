package main

import (
	"fmt"
	"errors"
	"os"
)

func doubleEven(i int) (int, error) { // 偶数なら2倍して戻す //liststart1
	if i % 2 != 0 { // 偶数ではないのでエラーを戻す
		return 0, errors.New("処理対象は偶数のみです")
	}
	return i * 2, nil  // 2倍して戻す
} 

func main() {
	i := 19
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err) // 処理対象は偶数のみです
		os.Exit(1)
	}
	fmt.Printf("%dの2倍: %d\n", i, double)
} //listend1
