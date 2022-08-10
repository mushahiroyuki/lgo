package calc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type stringStack []string

func (ss stringStack) Peek() string {
	if ss == nil || len(ss) == 0 {
		return ""
	}
	return ss[0]
}

func (ss stringStack) Pop() (stringStack, string) {
	if ss == nil || len(ss) == 0 {
		return ss, ""
	}
	result := ss[0]
	ss = ss[1:]
	return ss, result
}

func (ss stringStack) Push(s string) stringStack {
	ss = append(append(stringStack(nil), s), ss...)
	return ss
}

type floatStack []float64

func (fs floatStack) Peek() float64 {
	if fs == nil || len(fs) == 0 {
		return 0
	}
	return fs[0]
}

func (fs floatStack) Pop() (floatStack, float64) {
	if fs == nil || len(fs) == 0 {
		return fs, 0
	}
	result := fs[0]
	fs = fs[1:]
	return fs, result
}

func (fs floatStack) Push(f float64) floatStack {
	fs = append(append(floatStack(nil), f), fs...)
	return fs
}

func shouldEvaluate(newOp string, topOp string) bool {
	if topOp == "" || topOp == "(" {
		return false
	}

	// with 4 standard operators, the only time you don't evaluate is
	// when the new operator is a * or / and the top operator is a + or -
	// topOp     	newOp     	shouldEvaluate
	// -----     	----- 	    --------------
	// +, -       	+, -      	true
	// *, /       	+, -      	true
	// +, -       	*, /      	false
	// *, /       	*, /      	true
	if (topOp == "+" || topOp == "-") && (newOp == "*" || newOp == "/") {
		return false
	}
	return true
}

func eval(operators stringStack, numbers floatStack) (stringStack, floatStack, error) {
	var op string
	var first, second float64
	operators, op = operators.Pop()
	numbers, second = numbers.Pop()
	numbers, first = numbers.Pop()
	var result float64
	switch op {
	case "+":
		result = first + second
	case "-":
		result = first - second
	case "*":
		result = first * second
	case "/":
		result = first / second
	default:
		return nil, nil, errors.New("invalid operator: " + op)
	}
	numbers = numbers.Push(result)
	return operators, numbers, nil
}

func Process(expression string) (float64, error) {
	tokens := strings.Fields(expression)
	var operators stringStack
	var numbers floatStack
	for _, t := range tokens {
		switch t {
		case "+", "-", "/", "*":
			for shouldEvaluate(t, operators.Peek()) {
				var err error
				operators, numbers, err = eval(operators, numbers)
				if err != nil {
					return 0, fmt.Errorf("invalid expression %s: %w", expression, err)
				}
			}
			operators = operators.Push(t)
		case "(":
			operators = operators.Push(t)
		case ")":
			for op := operators.Peek(); op != "("; op = operators.Peek() {
				var err error
				operators, numbers, err = eval(operators, numbers)
				if err != nil {
					return 0, fmt.Errorf("invalid expression %s: %w", expression, err)
				}
			}
			operators, _ = operators.Pop()
		default:
			num, err := strconv.ParseFloat(t, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid expression %s: %w", expression, err)
			}
			numbers = numbers.Push(num)
		}
	}
	for op := operators.Peek(); op != ""; op = operators.Peek() {
		var err error
		operators, numbers, err = eval(operators, numbers)
		if err != nil {
			return 0, fmt.Errorf("invalid expression %s: %w", expression, err)
		}
	}
	numbers, out := numbers.Pop()
	if len(numbers) != 0 {
		return 0, fmt.Errorf("invalid expression %s", expression)
	}
	return out, nil
}
