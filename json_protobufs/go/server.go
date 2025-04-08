package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) < 4 {
		panic("insufficient length")
	}

	go func() {
		compare()
		return
		mux := http.NewServeMux()
		mux.HandleFunc("/default", func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			var res Small
			err = json.Unmarshal(b, &res)
			if err != nil {
				panic(err)
			}
			w.Write([]byte("end"))
		})
		mux.HandleFunc("/proto", func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			var res Small
			err = proto.Unmarshal(b, &res)
			if err != nil {
				panic(err)
			}
			w.Write([]byte("end"))
		})
		http.ListenAndServe(":3000", mux)
	}()
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)
	makeRequests(os.Args[1], count, os.Args[3])
}

func makeRequests(t string, count int, size string) {
	file, err := os.ReadFile(filepath.Join("..", "common", "json", fmt.Sprintf("%v.json", size)))
	if err != nil {
		panic(err)
	}

	var j Small
	err = json.Unmarshal(file, &j)
	if err != nil {
		panic(err)
	}

	if t == "1" {
		start := time.Now()
		for i := 0; i < count; i++ {
			m, err := json.Marshal(j)
			if err != nil {
				panic(err)
			}
			resp, err := http.Post(
				"http://localhost:3000/default",
				"application/json",
				bytes.NewBuffer(m),
			)
			if err != nil {
				panic(err)
			}
			resp.Body.Close()
		}
		fmt.Println(time.Since(start))
		return
	}

	start := time.Now()
	for i := 0; i < count; i++ {
		p, err := proto.Marshal(&j)
		if err != nil {
			panic(err)
		}
		resp, err := http.Post(
			"http://localhost:3000/proto",
			"application/octet-stream",
			bytes.NewBuffer(p),
		)
		if err != nil {
			panic(err)
		}
		resp.Body.Close()
	}
	fmt.Println(time.Since(start))
}
