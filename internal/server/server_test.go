package server_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jshawl/minimeter/internal/db"
	"github.com/jshawl/minimeter/internal/server"
)

func TestHandleGetApiMetrics(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/metrics", nil)
	w := httptest.NewRecorder()

	server.HandleGetApiMetrics(w, req)

	body := w.Body.String()

	if body != `{"ok": 1}` {
		t.Errorf("got %s, want {\"ok\": 1}", body)
	}
}

func TestHandlerPostApiMeasure_EnqueuesJob(t *testing.T) {
	jobs := make(chan db.Measurement, 1)
	w := httptest.NewRecorder()
	handler := server.HandlePostApiMeasure(jobs)
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/metrics",
		strings.NewReader(`{"name": "test_metric", "value": 42}`),
	)
	req.Header.Set("Content-Type", "application/json")
	handler(w, req)
	select {
	case job := <-jobs:
		if job.Name != "test_metric" || job.Value != 42 {
			t.Fatalf("unexpected job: %+v", job)
		}
	default:
		t.Fatal("expected job to be enqueued, but channel was empty")
	}
}
