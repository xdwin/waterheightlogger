package main

import (
	"fmt"
	"log"
	"net/http"

	db "./db"
)

func init() {
	fmt.Println("init main")
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprintf(w, path)
	db.Save()
}

func handleRoute() {
	http.HandleFunc("/", handler)
}

func startServer() {
	fmt.Println("Starting web server at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	handleRoute()
	startServer()
}
