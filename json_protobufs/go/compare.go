package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <level>")
	}
	level := os.Args[1]

	// Читаем JSON файл
	jsonFilePath := filepath.Join("..", "common", "json", fmt.Sprintf("%s.json", level))
	jsonDataBytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	// Парсим JSON
	var jsonData Large // Структура из Protobuf
	err = json.Unmarshal(jsonDataBytes, &jsonData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Кодируем в Protobuf
	protoData, err := proto.Marshal(&jsonData)
	if err != nil {
		log.Fatalf("Failed to marshal Protobuf: %v", err)
	}

	// Benchmark JSON Serialize
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_, _ = json.Marshal(jsonData)
	}
	fmt.Printf("JSON Serialize: %v\n", time.Since(start))

	// Benchmark Protobuf Serialize
	start = time.Now()
	for i := 0; i < 10000; i++ {
		_, _ = proto.Marshal(&jsonData)
	}
	fmt.Printf("Protobuf Serialize: %v\n", time.Since(start))

	// Benchmark JSON Parse
	start = time.Now()
	for i := 0; i < 10000; i++ {
		var tmp Large
		_ = json.Unmarshal(jsonDataBytes, &tmp)
	}
	fmt.Printf("JSON Parse: %v\n", time.Since(start))

	// Benchmark Protobuf Parse
	start = time.Now()
	for i := 0; i < 10000; i++ {
		var tmp Large
		_ = proto.Unmarshal(protoData, &tmp)
	}
	fmt.Printf("Protobuf Parse: %v\n", time.Since(start))
}
