package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Model struct {
	db *sql.DB
}

func NewDB(path string) (Model, error) {
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return Model{db: database}, err
	}

	sqlStmt := `
		create table if not exists measurements (
			id integer not null primary key,
			timestamp integer not null,
			name text not null,
			value real not null default 1
		);
		CREATE INDEX if not exists idx_measurements_name_timestamp ON measurements(name, timestamp DESC);
	`

	_, err = database.ExecContext(context.Background(), sqlStmt)

	return Model{db: database}, err
}

func (model Model) Measure(name string, value float64) (int64, error) {
	tx, err := model.db.BeginTx(
		context.Background(),
		&sql.TxOptions{ReadOnly: false, Isolation: 0},
	)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		context.Background(),
		"insert into measurements (name, value, timestamp) values (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	timestamp := time.Now().UnixMilli()
	result, err := stmt.ExecContext(context.Background(), name, value, timestamp)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

type Measurement struct {
	ID        int64
	Name      string
	Value     float64
	Timestamp int64
}

type MetricSummary struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
	Sum   int64  `json:"sum"`
	Avg   int64  `json:"avg"`
	Min   int64  `json:"min"`
	Max   int64  `json:"max"`
}

func (model Model) StartMeasurementWorker(jobs <-chan Measurement) {
	go func() {
		for job := range jobs {
			_, err := model.Measure(job.Name, job.Value)
			if err != nil {
				log.Printf("metric insert failed: %v", err)
			}
		}
	}()
}

func (model Model) GetMeasurements() ([]MetricSummary, error) {
	tx, err := model.db.BeginTx(
		context.Background(),
		&sql.TxOptions{ReadOnly: true},
	)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := `SELECT 
				name,
				COUNT(*),
				SUM(value),
				AVG(value),
				MIN(value),
				MAX(value)
				FROM measurements
				WHERE timestamp > (?)
				GROUP BY name`

	stmt, err := tx.PrepareContext(
		context.Background(),
		query,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	oneWeekAgo := time.Now().Add(-7 * 24 * time.Hour).UnixMilli()
	rows, err := stmt.QueryContext(context.Background(), oneWeekAgo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []MetricSummary{}

	for rows.Next() {
		var m MetricSummary
		if err := rows.Scan(&m.Name, &m.Count, &m.Sum, &m.Avg, &m.Min, &m.Max); err != nil {
			return nil, err
		}
		results = append(results, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return results, nil
}

func (db Model) Close() {
	db.db.Close()
}
