package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jshawl/minimeter/internal/db"
)

func HandleGetApiMetrics(db db.Model) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		metrics, _ := db.GetMeasurements()
		json.NewEncoder(w).Encode(metrics)
	}
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

func NewServer() (http.Handler, db.Model) {
	mux := http.NewServeMux()

	jobs := make(chan db.Measurement, 50_000)

	db, err := db.NewDB(os.Getenv("DB_PATH") + "minimeter.db")
	if err != nil {
		log.Fatal(err)
	}
	db.StartMeasurementWorker(jobs)
	mux.HandleFunc("/api/metrics", HandleGetApiMetrics(db))
	mux.HandleFunc("/api/measure", HandlePostApiMeasure(jobs))
	return mux, db
}

func ListenAndServe() {
	handler, db := NewServer()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()
	log.Println("Server running on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
