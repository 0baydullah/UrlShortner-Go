package main

import (
	"UrlShortner/config"
	"UrlShortner/handlers"
	"log"
	"net/http"
)

func main() {
	config.InitDB()

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Listening on port 8080")

	http.HandleFunc("/shorten", handlers.ShortenUrlHandler)
	http.HandleFunc("/", handlers.RedirectHandler)
}
