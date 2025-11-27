package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
