// ex1101b.go にUTF8の例を書きました

package main

import (
	"io"
	"fmt"
	"strings"
)

//liststart1
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf) // io.Readerの唯一のメソッド
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
//listend1

func countLettersInString() error {
	s := "The quick brown fox jumped over the lazy dog" //liststart2
	sr := strings.NewReader(s)  // 文字列sからio.Readerを作る
	counts, err := countLetters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts) //listend2
	return nil
}

func main() {
	countLettersInString()
}
