package main

import(
	"fmt"
	"time"
)

func main() {
//liststart1
	d := 2 * time.Hour + 30 * time.Minute + 45 * time.Second  // d の型はtime.Duration
	fmt.Println(d)  // 2h30m45s
//listend1
}
