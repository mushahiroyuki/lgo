package table

import "testing"

func TestDoMath(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"addition", 2, 2, "+", 4, ""},
		{"subtraction", 2, 2, "-", 0, ""},
		{"multiplication", 2, 2, "*", 4, ""},
		{"division", 2, 2, "/", 1, ""},
		{"bad_division", 2, 0, "/", 0, `0による除算`},
//liststart
		{"bad_op", 2, 2, "?", 0, `未知の演算子 ?`},
//listend		
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("期待する結果 %d, 得られた結果 %d", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("期待するエラーメッセージ `%s`, 得られたエラーメッセージ `%s`",
					d.errMsg, errMsg)
			}
		})
	}
}
