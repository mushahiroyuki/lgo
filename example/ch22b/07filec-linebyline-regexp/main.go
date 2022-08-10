// 1行ずつ正規表現を使って処理。regexp.FindAllStringSubmatchを利用
package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

var re1 = regexp.MustCompile(`^([^|]+)\|([^|]+)$`)
var re2 = regexp.MustCompile(`[,、，] *`)

func main(){
	data, _ := os.Open("testdata/dict2.txt")
	defer data.Close()
	
	scanner := bufio.NewScanner(data)
	for scanner.Scan(){
		s := scanner.Text()
		if s == "" || s[0:1] == "#" {
			continue  // この行はスキップ（コメントあるいは空白行）
		}
		matchResults := re1.FindAllStringSubmatch(s, -1)  // 結果の型は[][]string   // re1: `^([^|]+)\|([^|]+)$`
		// matchResults[0] -- マッチする全体
		// fmt.Println(matchResults[0])
		// fmt.Println(matchResults[0][1]) // 1つ目のマッチ
		// fmt.Println(matchResults[0][2]) // 2つ目のマッチ
		if len(matchResults[0])!=3 {
			continue // この行はスキップ。形式がおかしい
		}
		eng := matchResults[0][1]
		jpn := matchResults[0][2]
		translations := re2.Split(jpn, -1)  // `[,、，] *`
		for _, translation := range translations {
			fmt.Printf("%s|%s\n", eng, translation);
		}
	}
}
