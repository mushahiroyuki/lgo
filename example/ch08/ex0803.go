package main

import (
	"fmt"
	"os"
)

func doubleEven(i int) (int, error) { //liststart1
	if i % 2 != 0 {
		return 0, fmt.Errorf("%dは偶数ではありません", i)
	}
	return i * 2, nil
} 

func main() {
	i := 19
	double, err := doubleEven(i)
	if err != nil {
		fmt.Println(err)  // 19は偶数ではありません
		os.Exit(1)
	}
	fmt.Printf("%dの2倍: %d\n", i, double)
} //listend1
