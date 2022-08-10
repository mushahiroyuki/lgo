package main

import(
	"fmt"
	"errors"
	"time"
	"math/rand"
)

func timeLimit() (int, error) {  //liststart1
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	
	select {  // doneが閉じられるか
	case <-done:
		return result, err
	case <-time.After(2 * time.Second): // 2秒経過してしまうとこちら
		return 0, errors.New("タイムアウトしました")
	}
} //listend1

func doSomeWork() (int, error) {
	rand.Seed(time.Now().UnixNano()) // シードの指定
	n := rand.Intn(4)
	fmt.Println("n:", n)
	time.Sleep(time.Duration(n)*time.Second)
	result := 33
	return result, nil
}

func main() {
	result, err := timeLimit()

	if (err != nil) {
		fmt.Println(err)
	}	else {
		fmt.Printf("結果: %d\n", result)
	}
}
