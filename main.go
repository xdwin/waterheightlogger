package main

import (
	"fmt"
	"log"
	"net/http"

	heightLogger "github.com/xdwin/waterheightlogger/controller/heightlogger"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Fprintf(w, path)
}

func handleRoute() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/log/", heightLogger.Handler)
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
