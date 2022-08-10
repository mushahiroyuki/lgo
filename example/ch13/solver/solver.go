package solver

import (
	"context"
_	"errors"
	"io"
)

type Processor struct { //liststart1
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
} //listend1

func (p Processor) ProcessExpression(ctx context.Context, r io.Reader) (float64, error) { //liststart2
	curExpression, err := readToNewLine(r)
	if err != nil {
		return 0, err
	}
	if len(curExpression) == 0 {
		return 0, nil
//		return 0, errors.New("数式がありません")
	}
	answer, err := p.Solver.Resolve(ctx, curExpression)
	return answer, err
} //listend2

func readToNewLine(r io.Reader) (string, error) {
	var out []byte
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return string(out), nil
			}
		}
		if b[0] == '\n' {
			break
		}
		out = append(out, b[0])
	}
	return string(out), nil
}
