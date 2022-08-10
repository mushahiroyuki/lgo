package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"regexp"
)

var re1 = regexp.MustCompile(`[,，、] *`)

func main(){
	data, _ := os.Open("testdata/dict2.txt")
	defer data.Close()
	
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" || s[0:1] == "#" {
			continue  // この行はスキップ（コメントあるいは空白行）
		}
		columns := strings.Split(s, "|")
		if len(columns)!=2 {
			continue // この行はスキップ。形式がおかしい
		}
		eng := columns[0]
		jpn := columns[1]
		translations := re1.Split(jpn, -1)
		for _, translation := range translations {
			fmt.Printf("%s|%s\n", eng, translation);
		}
	}
}
