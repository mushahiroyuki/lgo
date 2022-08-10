package main

import (
	"fmt"
	"errors"
	"os"
)

//liststart
func divAndRemainder(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("0で割ることはできません")
		return 
	}
	result, remainder = numerator/denominator, numerator%denominator
	return
}
//listend

func callDivAndRemainder(numerator int, denominator int) {
	x, y, z := divAndRemainder(numerator, denominator)
	if z!=nil {
		fmt.Print(numerator, "÷", denominator, "：")
		fmt.Println(z)
		os.Exit(1)
	}
	fmt.Print(numerator, "÷", denominator, " = ", x, "余り" , y, "\n")
}

func main() {
	callDivAndRemainder(5,2)
	callDivAndRemainder(10,0)
}

