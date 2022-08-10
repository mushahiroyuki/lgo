package main

import (
	"fmt"
	"errors"
)

//liststart
func divAndRemainder(numerator int, denominator int) (result int, remainder int,
	err error) {	
	result, remainder = 20, 30  // 適当な値を代入
	if denominator == 0 {
		return numerator, denominator, errors.New("0で割ることはできません")
	}
	return numerator/denominator, numerator%denominator, nil
}

func main() {
	rs, rm, _ := divAndRemainder(5,2)
	fmt.Println(rs, rm)
}
