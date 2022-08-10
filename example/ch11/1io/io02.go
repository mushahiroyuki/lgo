// ex1102b.go にUTF8の例を書きました
package main

import (
	"io"
	"fmt"
	"os"
	"compress/gzip"
)


func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}


func buildGZipReader(fileName string) (*gzip.Reader, func(), error) { //liststart1
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
} //listend1

func countLettersInGzipFile() error {
	r, closer, err := buildGZipReader("my_data.txt.gz") //liststart2
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts) //listend2
	return nil
}

func main() {
	countLettersInGzipFile()
}
