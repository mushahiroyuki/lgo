package main

import (
	"archive/zip" // zip.ErrFormat が定義されている
	"bytes"
	"fmt"
)

func main() {
	data := []byte("This is not a zip file")
	nonZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(nonZipFile, int64(len(data)))
	if err == zip.ErrFormat {
		fmt.Println("ZIP形式ではありません")
	}
}
