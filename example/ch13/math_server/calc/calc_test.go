package calc

import "testing"

func TestProcess(t *testing.T) {
	data := []struct {
		name   string
		in     string
		out    float64
		errMsg string
	}{
		{
			name:   "simple +",
			in:     "2 + 3",
			out:    5,
			errMsg: "",
		},
		{
			name:   "simple -",
			in:     "2 - 3",
			out:    -1,
			errMsg: "",
		},
		{
			name:   "simple +",
			in:     "2 * 3",
			out:    6,
			errMsg: "",
		},
		{
			name:   "simple /",
			in:     "2 / 4",
			out:    0.5,
			errMsg: "",
		},
		{
			name:   "complex",
			in:     "2 * ( 3 + 2 ) / 5",
			out:    2,
			errMsg: "",
		},
		{
			name:   "order of operations",
			in:     "2 + 3 * 2 / 5",
			out:    3.2,
			errMsg: "",
		},
		{
			name:   "bad",
			in:     "( 2 + 3 * 2 / 5",
			out:    0,
			errMsg: "invalid expression ( 2 + 3 * 2 / 5: invalid operator: (",
		},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			out, err := Process(d.in)
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if out != d.out {
				t.Error(out)
			}
			if errMsg != d.errMsg {
				t.Error(errMsg)
			}
		})
	}
}
