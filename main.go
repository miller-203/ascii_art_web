package main

import (
	"log"
	"net/http"

	"ascii-art-web/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", handlers.ServeTemplate)
	http.HandleFunc("/ascii-art", handlers.HandleAsciiArt)
	log.Print("Listening on :http://localhost:8082/")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
