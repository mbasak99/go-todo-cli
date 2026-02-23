package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_name TEXT NOT NULL UNIQUE,
		task_desc TEXT NOT NULL,
		due_date TEXT,
		completed INTEGER DEFAULT 0
	)`)

	return db
}

func TestPerfectDataInsertTodo(t *testing.T) {
	t.Log("this performs an insert with all the fields with valid data")
	db := setupTestDB(t)
	defer db.Close()

	todo := Todo{
		Name: "Purchase flour and eggs",
		Desc: "Go to NoFrills and buy a carton of eggs and bag of flour.",
		DueDate: sql.NullString{
			String: time.Time.Local(time.Now()).Format("2006-01-02"),
			Valid:  true,
		},
		Completed: true,
	}

	_, err := db.Exec(`INSERT INTO todos (task_name, task_desc, due_date, completed)
	VALUES (?,?,?,?)`, todo.Name, todo.Desc, todo.DueDate.String, todo.Completed)
	if err != nil {
		t.Fatalf("failed to insert valid perfect data: %v", err)
	}
}

func TestPartialDataInsertTodo(t *testing.T) {
	t.Log("this performs an insert with only the mandatory fields with valid data")
	db := setupTestDB(t)
	defer db.Close()

	todo := Todo{
		Name: "Purchase flour and eggs",
		Desc: "Go to NoFrills and buy a carton of eggs and bag of flour.",
	}

	_, err := db.Exec(`INSERT INTO todos (task_name, task_desc)
	VALUES (?,?)`, todo.Name, todo.Desc)
	if err != nil {
		t.Fatalf("failed to insert valid partial data: %v", err)
	}
}
