package main

import (
	"github.com/jshawl/minimeter/internal/server"
)

func main() {
	handler := server.NewServer()
	server.ListenAndServe(handler)
}
