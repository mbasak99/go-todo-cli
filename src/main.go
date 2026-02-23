package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID        int
	Name      string
	Desc      string
	DueDate   sql.NullString
	Completed bool
}

func printTable(db *sql.DB) {
	rows, err := db.Query(`SELECT id, task_name, task_desc, due_date, completed FROM todos`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Todo
		rows.Scan(&t.ID, &t.Name, &t.Name, &t.DueDate, &t.Completed)

		// check for valid date
		if t.DueDate.Valid {
			fmt.Printf("%d %s %s %s %t\n", t.ID, t.Name, t.Desc, t.DueDate.String, t.Completed)
		} else {
			fmt.Printf("%d %s %s %t\n", t.ID, t.Name, t.Desc, t.Completed)
		}
	}
}

func main() {
	// set accepted flags
	todoFlag := flag.String("todo", "", "Create a new todo task with the name provided or updates an existing todo task.")
	descFlag := flag.String("desc", "", "Describes the todo's task if new or updates an existing description.")
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

	// disregard todos that don't have a task name and/or description
	if *todoFlag == "" {
		log.Fatalln("Please use the -todo=\"some task name\"")
	} else if *descFlag == "" {
		log.Fatalln("Please use the -desc=\"some task desc\"")
	}

	fmt.Println(*todoFlag, *descFlag, date.Format("2006/01/02"), *doneFlag)

	// create persistent storage with SQLite
	db, err := sql.Open("sqlite3", "./todos.app")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// create table for todos
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task_name TEXT NOT NULL UNIQUE,
		task_desc TEXT NOT NULL,
		due_date TEXT,
		completed INTEGER DEFAULT 0
	)`)

	// insert or update logic
	fmt.Println("Inserting todo to table...")
	_, err = db.Exec(`
	INSERT INTO todos (task_name, task_desc, due_date, completed)
	VALUES (?,?,?,?)
	ON CONFLICT(task_name) DO UPDATE SET
		task_desc = excluded.task_desc,
		due_date = excluded.due_date,
		completed = excluded.completed`,
		*todoFlag, *descFlag, date.Format("2006-01-02"), *doneFlag)

	// find the inserted row
	fmt.Println("Finding rows in todos table...")
	printTable(db)

	fmt.Println("Done.")
}
