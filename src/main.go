package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// set accepted flags
	todoFlag := flag.String("todo", "task", "Create a new todo task with the name provided or updates an existing todo task.")
	descFlag := flag.String("desc", "desciption", "Describes the todo's task if new or updates an existing description.")
	var date time.Time
	flag.Func("due", "Due date in YYYY-MM-DD format. If none provided there's no set due date.", func(s string) error {
		due, err := time.ParseInLocation("2006-01-02", s, time.Now().Local().Location())
		if err != nil {
			return err
		}
		date = due
		return nil
	})
	doneFlag := flag.Bool("completed", false, "By default creates an incomplete task.")

	flag.Parse()
	fmt.Println(*todoFlag, *descFlag, date.Format("2006/01/02"), *doneFlag)

	// create persistent storage with SQLite
	db, err := sql.Open("sqlite3", "./todos.app")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// create table for todos
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_name TEXT NOT NULL
		task_desc TEXT,
		due_date TEXT,
		completed INTEGER DEFAULT 0
	)`)
}
