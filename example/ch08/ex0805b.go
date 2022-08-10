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

func LoginAndGetData(uid, pwd, file string) ([]byte, error) { //liststart3
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:    InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
								// 認証情報が無効
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:    NotFound,
			Message: fmt.Sprintf("file %s not found", file), // ファイルが見つからない
		}
	}
	return data, nil
} //listend3


func GenerateError(flag bool) error { //liststart
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
} //listend


func login(uid, pwd string) error {
	return nil
}

func getData(file string) ([]byte, error) {
	var dummy []byte;
	return dummy, nil
}

