/*
$ go run ex0720.go

続いてブラウザで次のURLなどを表示（user_id= のあとに1, 2, 3などを指定）

http://localhost:8080/hello?user_id=1

*/

package main

import (
	"net/http"
	"errors"
	"fmt"
)


func LogOutput(message string) {
	fmt.Println(message)

}

type SimpleDataStore struct {
    userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
    name, ok := sds.userData[userID]
    return name, ok
}


func NewSimpleDataStore() SimpleDataStore {
    return SimpleDataStore{
        userData: map[string]string{
            "1": "Fred",
            "2": "Mary",
            "3": "Pat",
        },
    }
}

type DataStore interface {
    UserNameForID(userID string) (string, bool)
}

type Logger interface {
    Log(message string)
}

type LoggerAdapter func(message string)


func (lg LoggerAdapter) Log(message string) {
    lg(message)
}

type SimpleLogic struct {
    l  Logger
    ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
    sl.l.Log("SayHello(" + userID + ")")
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("不明なユーザー")
    }
    return name + "さん、こんにちは。", nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
    sl.l.Log("SayGoodbye(" + userID + ")")
    name, ok := sl.ds.UserNameForID(userID)
    if !ok {
        return "", errors.New("不明なユーザー")
    }
    return name + "さんさようなら。", nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
    return SimpleLogic{
        l:    l,
        ds: ds,
    }
}


type Logic interface {
    SayHello(userID string) (string, error)
}

type Controller struct {
    l     Logger
    logic Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
    c.l.Log("SayHello内: ")
    userID := r.URL.Query().Get("user_id")
    message, err := c.logic.SayHello(userID)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
        return
    }
    w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
    return Controller{
        l:     l,
        logic: logic,
    }
}

func main() {
    l := LoggerAdapter(LogOutput)
    ds := NewSimpleDataStore()
    logic := NewSimpleLogic(l, ds)
    c := NewController(l, logic)
    http.HandleFunc("/hello", c.SayHello)
    http.ListenAndServe(":8080", nil)
}






