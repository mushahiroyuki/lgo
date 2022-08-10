package main  //liststart1

import (
	"fmt"
	"time"
)

type Counter struct { // Counter型を定義
	total             int  // 合計
	lastUpdated time.Time  // 最終更新時刻
}

func (c *Counter) Increment() { // ポインタレシーバ（cはポインタ）
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // 値レシーバ（cのコピーが渡される）
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
} //listend1


func main() { //liststart2
	var c Counter
	fmt.Println(c.String())
	(&c).Increment()
	fmt.Println(c.String()) 
} //listend2
