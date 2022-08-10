package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]  // [1 2]  （xの先頭から2つ。つまりx[0]とx[1]）
	fmt.Println(cap(x), cap(y)) // 4  4
	y = append(y, 30)
	fmt.Println("x:", x) // [1 2 30 4]
	fmt.Println("y:", y) // [1 2 30]
}

/*
    x[0]  x[1]  x[2]   x[3]
    y[0]  y[1]
-----------------------------
    1     2     3     4     ← x := []int{1, 2, 3, 4}
    1     2    30     4     ← y = append(y, 30)
  {-------xの範囲---------}  → [1 2 30 4]
  {----yの範囲----}          → [1 2 30]
  y[0]  y[1]  y[2]         
*/
