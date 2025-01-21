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

var numOfRuns = 10000

func compare() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <level>")
	}
	level := os.Args[1]

	jsonFilePath := filepath.Join("..", "common", "json", fmt.Sprintf("%s.json", level))
	jsonDataBytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var jsonData Large
	err = json.Unmarshal(jsonDataBytes, &jsonData)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	protoData, err := proto.Marshal(&jsonData)
	if err != nil {
		log.Fatalf("Failed to marshal Protobuf: %v", err)
	}

	start := time.Now()
	for i := 0; i < numOfRuns; i++ {
		_, _ = json.Marshal(jsonData)
	}
	fmt.Printf("JSON Serialize: %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < numOfRuns; i++ {
		_, _ = proto.Marshal(&jsonData)
	}
	fmt.Printf("Protobuf Serialize: %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < numOfRuns; i++ {
		var tmp Large
		_ = json.Unmarshal(jsonDataBytes, &tmp)
	}
	fmt.Printf("JSON Parse: %v\n", time.Since(start))

	start = time.Now()
	for i := 0; i < numOfRuns; i++ {
		var tmp Large
		_ = proto.Unmarshal(protoData, &tmp)
	}
	fmt.Printf("Protobuf Parse: %v\n", time.Since(start))
}
