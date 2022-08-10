package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
//liststart1	
	const data = `
		{"name": "フレッド", "age": 40}
		{"name": "メアリ", "age": 21}
		{"name": "パッド", "age": 30}
	` //listend1
	var t struct {  //liststart2
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b) // bにデコードした結果を入れる

	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		} //listend2
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("------")	
	out := b.String()
	fmt.Println(out)
}
