package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	port := 8080

	h1 := func(w http.ResponseWriter, r *http.Request) {
		log.Printf(">> New request - %s %s\n", r.Method, r.URL)
		io.WriteString(w, "Hello!\n")
	}

	http.HandleFunc("/", h1)

	log.Printf(">> Server runnning at :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
