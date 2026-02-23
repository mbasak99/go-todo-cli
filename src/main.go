package main

import (
	"flag"
	"fmt"
	"time"
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
}
