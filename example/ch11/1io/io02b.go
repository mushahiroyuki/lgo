// ex1102.goをUTF8の文字についてカウントするようにしたものです
package main

import (
	"io"
	"fmt"
	"os"
	"compress/gzip"
)


func countUtf8Letters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		runes := []rune(string(buf[:n]))  // stringを経由せずに変換可能？
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

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func xxx() error {
	r, closer, err := buildGZipReader("my_datax.txt.gz")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countUtf8Letters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func main() {
	xxx()
}
