package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jshawl/minimeter/internal/db"
)

func HandleGetApiMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"ok\": 1}")
}

func HandlePostApiMeasure(jobs chan<- db.Measurement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			fmt.Fprint(w, `{"error": "method not allowed"}`)
			return
		}
		var measurement db.Measurement

		err := json.NewDecoder(r.Body).Decode(&measurement)

		if err != nil {
			log.Println(err)
			fmt.Fprint(w, `{"ok":1}`)
			return
		}

		select {
		case jobs <- measurement:
			// queued successfully
		default:
			// channel full
			log.Println("metric dropped: queue full")
		}
		fmt.Fprint(w, `{"ok":1}`)
	}
}

func NewServer() http.Handler {
	mux := http.NewServeMux()

	jobs := make(chan db.Measurement, 50_000)

	db, err := db.NewDB("minimeter.db")
	if err != nil {
		log.Fatal(err)
	}
	db.StartMeasurementWorker(jobs)
	mux.HandleFunc("/api/metrics", HandleGetApiMetrics)
	mux.HandleFunc("/api/measure", HandlePostApiMeasure(jobs))

	return mux
}

func ListenAndServe(handler http.Handler) {
	port := ":8080"
	log.Printf("Starting server on %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
