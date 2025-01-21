package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"
)

var jsonPool = sync.Pool{
	New: func() interface{} {
		return new(Small)
	},
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func main2() {
	if len(os.Args) < 4 {
		panic("insufficient arguments")
	}

	// Запускаем сервер ОТДЕЛЬНО
	if os.Args[1] == "server" {
		startServer()
		return
	}

	// Ждем, чтобы сервер точно запустился
	time.Sleep(1 * time.Second)

	// Запускаем клиент для теста
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	makeRequests(os.Args[1], count, os.Args[3])
}

// === СЕРВЕР ===
func startServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/default", handleJSONRequest)
	mux.HandleFunc("/proto", handleProtoRequest)

	server := &http.Server{
		Addr:         ":3000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	fmt.Println("Server started on port 3000")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	obj := jsonPool.Get().(*Small)
	defer jsonPool.Put(obj)

	buf := bufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bufPool.Put(buf)

	_, err := io.Copy(buf, r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(buf.Bytes(), obj); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	io.WriteString(w, "end")
}

func handleProtoRequest(w http.ResponseWriter, r *http.Request) {
	obj := jsonPool.Get().(*Small)
	defer jsonPool.Put(obj)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusInternalServerError)
		return
	}

	if err := proto.Unmarshal(body, obj); err != nil {
		http.Error(w, "invalid protobuf", http.StatusBadRequest)
		return
	}
	io.WriteString(w, "end")
}

// === КЛИЕНТ ===
func makeRequests2(reqType string, count int, size string) {
	file, err := os.ReadFile(fmt.Sprintf("../common/json/%v.json", size))
	if err != nil {
		panic(err)
	}

	var j Small
	if err := json.Unmarshal(file, &j); err != nil {
		panic(err)
	}

	client := &http.Client{}
	var wg sync.WaitGroup
	reqChan := make(chan struct{}, 100) // Ограничение 100 запросов одновременно

	start := time.Now()
	for i := 0; i < count; i++ {
		wg.Add(1)
		reqChan <- struct{}{} // Блокируем, если 100 запросов уже идут
		go func() {
			defer wg.Done()
			defer func() { <-reqChan }() // Освобождаем слот

			if reqType == "1" {
				sendJSONRequest(client, j)
			} else {
				sendProtoRequest(client, j)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Total time:", time.Since(start))
}

func sendJSONRequest(client *http.Client, data Small) {
	m, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/default", bytes.NewBuffer(m))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}

func sendProtoRequest(client *http.Client, data Small) {
	p, err := proto.Marshal(&data)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/proto", bytes.NewBuffer(p))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
}
