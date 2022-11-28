package io

import (
	"encoding/json"
	"log"
	"os"
)

func WriteToJsonFile(data interface{}, path string) {
	file, _ := json.MarshalIndent(data, "", " ")
	err := os.WriteFile(path, file, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
