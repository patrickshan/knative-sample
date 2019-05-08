package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("application received a request")
	target, defined := os.LookupEnv("TARGET")
	if !defined {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s! app version=0.2", target)
}
