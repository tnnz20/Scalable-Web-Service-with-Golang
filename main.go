package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tnnz20/Scalable-Web-Service-with-Golang/internal/handler"
	"github.com/tnnz20/Scalable-Web-Service-with-Golang/pkg/helpers"
)

func main() {
	// Update data every 5 seconds
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for range ticker.C {
			element := helpers.UpdateElement(1, 100)
			helpers.WriteJSON(element)
			log.Println("Element updated")
		}
	}()

	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
	http.HandleFunc("/", handler.RenderHTML)

	// Serve the web service
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
