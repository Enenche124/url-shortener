package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/Enenche124/url-shortener/internal/handlers"
	"github.com/Enenche124/url-shortener/internal/service"
	"github.com/Enenche124/url-shortener/internal/storage"
)

func main() {
	fmt.Println("🚀 URL Shortener is starting on http://localhost:8080")

	store := storage.NewURLStore()
	urlService := &service.URLService{Store: store}

	http.HandleFunc("/api/shorten", handler.MakeShortenHandler(urlService))
	http.HandleFunc("/", handler.MakeRedirectHandler(urlService))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
