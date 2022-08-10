package text

import (
	// 	"io/ioutil"
	"os"
	"unicode/utf8"
)

func CountCharacters(fileName string) (int, error) {
	data, err := os.ReadFile(fileName)
	// data, err := ioutil.ReadFile(fileName) 
	if err != nil {
		return 0, err
	}
	return utf8.RuneCount(data), nil
}
