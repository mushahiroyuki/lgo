package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total             int
	lastUpdated time.Time
}

func (c *Counter) Increment() { // cのアドレスが渡される
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string { // cのコピーが渡される
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) { // mainのcのコピー
	(&c).Increment()  // mainのcのコピーに対してIncrementが行われる
	fmt.Println("NG:", c.String())
}

func doUpdateRight(c *Counter) { // 正しい
	c.Increment()  // mainのcに対してIncrementが行われる
	fmt.Println("OK:", c.String())
}

func main() {
	var c Counter
	doUpdateWrong(c)
	fmt.Println("main:", c.String())
	doUpdateRight(&c)
	fmt.Println("main:", c.String())
}

