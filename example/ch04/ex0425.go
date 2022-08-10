package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().Unix()) // シードの指定 //liststart
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("少し小さすぎます:", n)
	case n > 5:
		fmt.Println("大きすぎます:", n)
	default:
		fmt.Println("いい感じの数字です:", n)
	} //listend
}
