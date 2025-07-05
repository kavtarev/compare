package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

func JsonHandler(name string) func(w http.ResponseWriter, req *http.Request) {
	file, err := os.ReadFile("../common/json/" + name + ".json")
	if err != nil {
		panic("no file founded")
	}

	var myMap map[string]any

	jErr := json.Unmarshal(file, &myMap)
	if jErr != nil {
		panic(jErr)
	}

	return func(w http.ResponseWriter, req *http.Request) {
		mar, err := json.Marshal(myMap)
		if err != nil {
			panic("cant stringify")
		}

		w.Write(mar)
	}
}
