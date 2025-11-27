package server

import (
	"fmt"
	"log"
	"net/http"
)

func HandleGetApiMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"ok\": 1}")
}

func NewServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/metrics", HandleGetApiMetrics)

	return mux
}

func ListenAndServe(handler http.Handler) {
	port := ":8080"
	log.Printf("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
