package main

import (
	"io"
	"log"
	"os"
)

func getFile(name string) (*os.File, func(), error) { //liststart1
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, err
} //listend1

func main() {
	if len(os.Args) < 2 {
		log.Fatal("ファイルが指定されていません")
	}
	f, closer, err := getFile(os.Args[1]) //liststart2
	if err != nil {
		log.Fatal(err)
	}
	defer closer() //listend2
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
