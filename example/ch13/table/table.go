package table

import (
	"errors"
	"fmt"
)

func DoMath(num1, num2 int, op string) (int, error) { //liststart
	switch op {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 + num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("0による除算")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("未知の演算子 %s", op)
	}
} //listend

