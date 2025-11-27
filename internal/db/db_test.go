package db_test

import (
	"os"
	"testing"

	"github.com/jshawl/minimeter/internal/db"
)

func TestNewDB(t *testing.T) {
	path := t.TempDir() + "/test.db"
	db.NewDB(path)

	_, err := os.Stat(path)

	fileExists := err == nil

	if !fileExists {
		t.Fatal("expected db to be created")
	}
}

func TestMeasure(t *testing.T) {
	path := t.TempDir() + "/test.db"
	model, err := db.NewDB(path)
	if err != nil {
		t.Fatal(err)
	}

	id, err := model.Measure("example", 42)
	if id != 1 {
		t.Fatal("expected id to be 1")
	}
	if err != nil {
		t.Fatal(err)
	}

	results, err := model.GetMeasurements("example")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result got %d", len(results))
	}

	result := results[0]
	if result.ID != 1 {
		t.Fatalf("expected id to be 1 got %d", result.ID)
	}
	if result.Name != "example" {
		t.Fatalf("expected id to be example got %s", result.Name)
	}
	if result.Value != 42 {
		t.Fatalf("expected value to be 42 got %f", result.Value)
	}
}
