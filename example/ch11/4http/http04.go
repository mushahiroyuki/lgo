package main

import(
	"fmt"
	"net/http"
	"log"
	"time"
)


func main() {
	// http.HandleFuncは、特定のパターンを処理するためのハンドラ関数をDefaultServeMuxに登録する関数
	// http.HandlerFuncは型の名前
	// type HandlerFunc func(ResponseWriter, *Request)

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




	terribleSecurity := TerribleSecurityProvider("GOPHER") //liststart3

	mux.Handle("/hello", terribleSecurity(RequestTimer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello!\n"))
	})))) //listend3

	fmt.Printf("%s", "次でテスト:\ncurl http://localhost:8080/dog/greet\ncurl http://localhost:8080/person/greet\n")
	
	log.Fatal(http.ListenAndServe(":8080", mux))		

}


func RequestTimer(h http.Handler) http.Handler { //liststart2
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
			}
			h.ServeHTTP(w, r)
		})
	}
} //listend2

