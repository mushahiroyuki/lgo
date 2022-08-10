package adder_test //liststart

import (
	"testing"

	"test_examples/adder3"
)

func TestAddNumbers(t *testing.T) {
	result := adder.AddNumbers(2, 3)
	if result != 5 {
		t.Error("結果が正しくない: 想定される結果 5, 得られた結果", result)
	}
} //listend
