package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/printHi", printHi)
	router.HandleFunc("/printBye", printBye)

	http.ListenAndServe(":8080", router)
}

func printHi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Hi")
}

func printBye(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Bye")
}
