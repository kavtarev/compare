package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

func ClosureJson(name string) func(w http.ResponseWriter, req *http.Request) {
	file, err := os.ReadFile("../common/json/" + name + ".json")
	if err != nil {
		panic("no file found")
	}

	var myMap map[string]any

	jErr := json.Unmarshal(file, &myMap)
	if jErr != nil {
		panic(jErr)
	}

	res := func(w http.ResponseWriter, req *http.Request) {
		mar, err := json.Marshal(myMap)
		if err != nil {
			panic("cant stringify")
		}

		w.Write(mar)
	}

	return res
}
