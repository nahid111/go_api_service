package main

import (
	"log"
	"net/http"
)

func writeResponse(writer http.ResponseWriter, responseStrig string) {
	res := []byte(responseStrig)
	_, err := writer.Write(res)
	if err != nil {
		log.Println(err)
	}
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(w, "Hello World!")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	writeResponse(w, "OK!")
}

func main() {
	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on %s ...", addr)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal(err)
	}
}
