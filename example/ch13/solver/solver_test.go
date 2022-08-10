package solver

import (
	"context"
	"errors"
	"strings"
	"testing"
)

type MathSolverStub struct{} //liststart1

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10":
		return 0, errors.New("不正な数式: ( 2 + 2 * 10")
	}
	return 0, nil
} //listend1

func TestProcessor_ProcessExpressions(t *testing.T) { //liststart2
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
( 2 + 2 ) * 10
( 2 + 2 * 10`)
	data := []float64{22, 40, 0, 0}
	for _, d := range data {
		result, err := p.ProcessExpression(context.Background(), in)
		if err != nil {
			t.Error(err)
		}
		if result != d {
			t.Errorf("期待される結果 %f, 得られた結果 %f", d, result)
		}
	}
} //listend2
