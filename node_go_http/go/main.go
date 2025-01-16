package main

import (
	"compare/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getpid())

	if len(os.Args) == 1 {
		fmt.Println("should pass file type argument")
		return
	}

	http.HandleFunc("/json-stringify", handlers.JsonHandler(os.Args[1]))
	http.HandleFunc("/download-file", handlers.ReadFileChunkHandler(os.Args[1]))
	http.HandleFunc("/parse-xml", handlers.XmlHandler(os.Args[1]))

	http.ListenAndServe(":3000", nil)
}
