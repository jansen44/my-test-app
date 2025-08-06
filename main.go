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
		log.Printf(">> Nova request - %s %s\n", r.Method, r.URL)
		io.WriteString(w, "Oier!\n")
	}

	http.HandleFunc("/", h1)

	log.Printf(">> Server rodando na porta :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
