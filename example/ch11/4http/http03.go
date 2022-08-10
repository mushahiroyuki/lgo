package main

import(
	"fmt"
	"net/http"
	"log"
)


func main() {
	// http.HandleFuncは、特定のパターンを処理するためのハンドラ関数をDefaultServeMuxに登録する関数
	// Muxはmultiplexerの省略形。Goでは「パス」と「そのパスを処理するハンドラ」を管理する機構のこと
	// http.HandlerFuncは型の名前

	personMux := http.NewServeMux() //liststart1
	personMux.HandleFunc("/greet",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("こんにちは！\n"))
		})
	
	dogMux := http.NewServeMux()
	dogMux.HandleFunc("/greet",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("カワイイ子犬ちゃんだね！\n"))
		})

	mux := http.NewServeMux()
	mux.Handle("/person/", http.StripPrefix("/person", personMux))
	// "/person/" のパターンは personMux が処理する
	mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))
	// "/dog/" のパターンは dogMux が処理する //listend1
	fmt.Printf("%s", "次でテスト:\ncurl http://localhost:8080/dog/greet\ncurl http://localhost:8080/person/greet\n")

	log.Fatal(http.ListenAndServe(":8080", mux))		

}
