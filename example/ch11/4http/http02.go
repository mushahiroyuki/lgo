package main

import(
	"fmt"
	"net/http"
	"time"
)


/*
http.Serverの定義（https://pkg.go.dev/net/http#Server 参照）
	type Server struct {
		Addr string
		Handler Handler // ←ここにインタフェースhttp.Handlerを満たすHandlerを代入する
		TLSConfig *tls.Config
		ReadTimeout time.Duration
		ReadHeaderTimeout time.Duration
		WriteTimeout time.Duration
		IdleTimeout time.Duration
		MaxHeaderBytes int
		TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
		ConnState func(net.Conn, ConnState)
		ErrorLog *log.Logger
		BaseContext func(net.Listener) context.Context
		ConnContext func(ctx context.Context, c net.Conn) context.Context
	}


// https://pkg.go.dev/net/http#Handler参照
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}

*/

type HelloHandler struct{} //liststart1

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
} //listend1

func main() {
	s := http.Server{ //liststart2
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      HelloHandler{},
	}
	fmt.Println("ブラウザで http://localhost:8080/ を開いてください。")
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	} //listend2

	/*
     ブラウザで http://localhost:8080 を開く
     あるいは  curl http://localhost:8080
  */
}


	
