package main

import (
	"fmt"
)

type Status int  //liststart1

const (  // iotaについては「7.2.7 iotaと列挙型」参照
	InvalidLogin Status = iota + 1  // 不正なログイン
	NotFound  // 見つからない
) //listend1

type StatusErr struct {  //liststart2
	Status  Status // 状態
	Message string // メッセージ
}

func (se StatusErr) Error() string {
	return se.Message
} //listend2

func GenerateError(flag bool) error { //liststart3
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
} //listend3
