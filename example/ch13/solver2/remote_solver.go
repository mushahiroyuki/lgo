package solver

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type RemoteSolver struct { //liststart1
	MathServerURL string
	Client        *http.Client
}

func (rs RemoteSolver) Resolve(ctx context.Context, expression string) (float64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		rs.MathServerURL+"?expression="+url.QueryEscape(expression), nil)
	if err != nil {
		return 0, err
	}
	resp, err := rs.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(contents))
	}
	result, err := strconv.ParseFloat(string(contents), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
} //listend1

