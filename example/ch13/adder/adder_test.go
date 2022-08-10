package adder

import "testing"

func Test_addNumbers(t *testing.T) { //liststart
	result := addNumbers(2, 3)
	if result != 5 {
		t.Error("結果が違う: 期待する結果 5, 得られた結果", result)
	}
} //listend

// t.Errorf("結果が違う: 期待する結果 %d, 得られた結果 %d", 5, result)
// t.Fatal("結果が違う: 期待する結果 5, 得られた結果", result)
// t.Fatalf("結果が違う: 期待する結果 %d, 得られた結果 %d", 5, result)

