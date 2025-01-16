package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ClosureReadFileHandlerChunk(name string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const size = 50000

		f, err := os.Open("../common/files/" + name + ".txt")
		if err != nil {
			panic("cant open file")
		}
		defer f.Close()

		bytes := make([]byte, size)

		for {
			bytesRead, err := f.Read(bytes)

			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Println(err)
			}

			w.Write(bytes[:bytesRead])
		}
	}
}
