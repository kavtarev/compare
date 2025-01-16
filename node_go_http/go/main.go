package main

import (
	"compare/handlers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getpid())

	http.HandleFunc("/json-stringify", handlers.ClosureJson(os.Args[1]))
	http.HandleFunc("/download-file", handlers.ClosureReadFileHandlerChunk(os.Args[1]))
	http.HandleFunc("/parse-xml", handlers.ClosureXmlHandler(os.Args[1]))

	http.ListenAndServe(":3000", nil)
}
