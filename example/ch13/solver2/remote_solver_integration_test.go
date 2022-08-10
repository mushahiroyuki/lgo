//go:build integration

package solver

import (
	"context"
	"net/http"
	"testing"
)

func TestRemoteSolver_ResolveIntegration(t *testing.T) {
	rs := RemoteSolver{
		MathServerURL: "http://localhost:8080",
		Client:        http.DefaultClient,
	}
	data := []struct {
		name       string
		expression string
		result     float64
		errMsg     string
	}{
		{"case1", "2 + 2 * 10", 22, ""},
		{"case2", "( 2 + 2 ) * 10", 40, ""},
		{"case3", "( 2 + 2 * 10", 0, "invalid expression: ( 2 + 2 * 10"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := rs.Resolve(context.Background(), d.expression)
			if result != d.result {
				t.Errorf("expected `%f`, got `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected error `%s`, got `%s`", d.errMsg, errMsg)
			}
		})
	}
}
