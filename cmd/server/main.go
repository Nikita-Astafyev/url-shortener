package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/Nikita-Astafyev/url-shortener/internal/config"
	"github.com/Nikita-Astafyev/url-shortener/internal/handler"
	"github.com/Nikita-Astafyev/url-shortener/internal/storage"
)

func main() {
	dbStorage, err := storage.NewPostgresStorage(config.GetDBConfig())
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbStorage.DB.Close()

	handler := handler.NewURLHandler(dbStorage)
	http.HandleFunc("/create", handler.CreateShortURL)
	http.HandleFunc("/r/", handler.Redirect)

	port := config.GetServerConfig().Port
	log.Printf("Server started at :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
