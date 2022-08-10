// ex1101.goをUTF8の文字についてカウントするようにしたものです
package main

import (
	"io"
	"fmt"
	"strings"
)

func countUtf8Letters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		runes := []rune(string(buf[:n]))
		for _, b := range runes {
			out[string(b)]++
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func xxx() error {
	s := "東京特許許可局から許可がおりた特許について特許庁に相談に行った。"
	sr := strings.NewReader(s)
	counts, err := countUtf8Letters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func main() {
	xxx()
}
