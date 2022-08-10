package main

import (
	"fmt"
	"errors"
	"os"
)

func main() { //liststart2
	numerator, denominator := 20, 3	
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%d÷%d: 商:%d, 余り: %d\n", numerator, denominator,
		remainder, mod)
} //listend2

//liststart1
func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	// 被除数と除数をもらって、商と余り、それにerrorを返す
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0") // 除数が0
	}

	return numerator / denominator, numerator % denominator, nil
}  //listend1
