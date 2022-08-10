package main

import "fmt"

func main() {
	x := make([]int, 0, 5)
	fmt.Println("x:", x)  // []
	x = append(x, 1, 2, 3, 4)
	fmt.Println("x:", x)  // [1 2 3 4]
	y := x[:2:2]
	fmt.Println("y:", y)  // [1 2]
	z := x[2:4:4]
	fmt.Println("z:", z)  // [3 4]
	fmt.Println("cap:", cap(x), cap(y), cap(z))  // 5 2 2
	fmt.Println("1-----")
	y = append(y, 30, 40, 50)
	fmt.Println("x:", x)  // [1 2 3 4]  （変わらない）
	fmt.Println("y:", y)  // [1 2 30 40 50]
	fmt.Println("z:", z)  // [3 4]  （変わらない）
	fmt.Println("cap:", cap(x), cap(y), cap(z))  // 5 6 2
	fmt.Println("2-----")
	x = append(x, 60)
	fmt.Println("x:", x)  // [1 2 3 4 60]
	fmt.Println("y:", y)  // [1 2 30 40 50] （変わらない）
	fmt.Println("z:", z)  // [3 4]  （変わらない）
	fmt.Println("cap:", cap(x), cap(y), cap(z))  // 5 6 2
	z = append(z, 70)
	fmt.Println("3-----")
	fmt.Println("x:", x)  // [1 2 3 4 60] （変わらない）
	fmt.Println("y:", y)  // [1 2 30 40 50] （変わらない）
	fmt.Println("z:", z)  // [3 4 70]
	fmt.Println("cap:", cap(x), cap(y), cap(z))  // 5 6 4
}
