package db_test

import (
	"os"
	"testing"

	"github.com/jshawl/minimeter/internal/db"
)

func TestNewDB(t *testing.T) {
	path := t.TempDir() + "/test.db"
	model, _ := db.NewDB(path)

	_, err := os.Stat(path)

	fileExists := err == nil

	if !fileExists {
		t.Fatal("expected db to be created")
	}
	defer model.Close()
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

	results, err := model.GetMeasurements()
	if err != nil {
		t.Fatal(err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result got %d", len(results))
	}

	result := results[0]
	if result.Name != "example" {
		t.Fatalf("expected id to be example got %s", result.Name)
	}
	if result.Count != 1 {
		t.Fatalf("expected count to be 1 got %d", result.Count)
	}
	if result.Sum != 42 {
		t.Fatalf("expected sum to be 42 got %d", result.Sum)
	}
	if result.Avg != 42 {
		t.Fatalf("expected avg to be 42 got %d", result.Avg)
	}
	if result.Min != 42 {
		t.Fatalf("expected min to be 42 got %d", result.Min)
	}
	if result.Max != 42 {
		t.Fatalf("expected max to be 42 got %d", result.Max)
	}
	defer model.Close()
}
