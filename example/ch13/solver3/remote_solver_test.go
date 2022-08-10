package solver

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {
	type info struct { //liststart1
		expression string
		code       int
		body       string
	}
	var io info //listend1
	server := httptest.NewServer( //liststart2
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			expression := req.URL.Query().Get("expression")
			if expression != io.expression {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("不正な数式: " + io.expression))
				return
			}
			rw.WriteHeader(io.code)
			rw.Write([]byte(io.body))
		}))
	defer server.Close()
	rs := RemoteSolver{
		MathServerURL: server.URL,
		Client:        server.Client(),
	} //listend2
	data := []struct { //liststart3
		name   string
		io     info
		result float64
		errMsg string
	}{
		{"case1", info{"2 + 2 * 10", http.StatusOK, "22"}, 22, ""},
		{"case2", info{"( 2 + 2 ) * 10", http.StatusOK, "40"}, 40, ""},
		{"case3", info{"( 2 + 2 * 10", http.StatusBadRequest,
			"不正な数式: ( 2 + 2 * 10"},
			0, "不正な数式: ( 2 + 2 * 10"},
	}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			io = d.io
			result, err := rs.Resolve(context.Background(), d.io.expression)
			if result != d.result {
				t.Errorf("予想されたIO `%f`, 得られた結果 `%f`", d.result, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("予想されたIOエラー `%s`, 得られたIOエラー `%s`", d.errMsg, errMsg)
			}
		})
	} //listend3
}
