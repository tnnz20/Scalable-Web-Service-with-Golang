package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/internal/domain"
)

var (
	PATH = "pkg/datasources/element.json"
)

func WriteJSON(element *domain.Element) {

	jsonData, err := json.Marshal(element)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	file, err := os.Create(PATH)
	if err != nil {
		log.Println("Error creating file:", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Println("Error writing JSON to file:", err)
		os.Exit(1)
	}
}

func LoadJSON() (*domain.Element, error) {
	jsonData, err := os.ReadFile(PATH)
	if err != nil {
		return nil, errors.New("error reading JSON file")
	}

	// Unmarshal the JSON data into a struct
	var data domain.Element
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, errors.New("error unmarshaling JSON")
	}

	return &data, nil
}
