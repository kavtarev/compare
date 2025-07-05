package sender

import (
	"fmt"
	"net/http"
)

type SenderServerOpts struct {
	Port string
}

func StartServer(opts SenderServerOpts) {
	mux := http.NewServeMux()
	mux.HandleFunc("/http", HttpHandler)

	err := http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("some")
	w.Write([]byte("hello"))
}
