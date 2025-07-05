package receiver

import (
	"fmt"
	"net/http"
)

type ReceiverServerOpts struct {
	Port string
}

func StartServerReceiver(opts ReceiverServerOpts) {
	mux := http.NewServeMux()
	mux.HandleFunc("/http", HttpHandler)

	fmt.Println("receiver before ListenAndServe")
	err := http.ListenAndServe(opts.Port, mux)
	if err != nil {
		panic(err)
	}
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("some")
	w.Write([]byte("hello"))
}
