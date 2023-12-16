package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{"Hello, World!"}
	js, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/message", messageHandler)
	http.ListenAndServe(":8080", nil)
}
