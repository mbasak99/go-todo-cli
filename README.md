# Go Todo CLI
Want to practice some of the Go concepts I've learned for a basic Todo CLI app.

# How To Run
`go run src/main.go -todo="task 1" -desc="My first task!"`

# Flags
- todo -> given name to task the user creates (string)
- desc -> given description to task the user creates (string)
- due -> when the task is due (string) (optional)
- completed -> whether the task is completed or not, when not provided automaitcally set as false (boolean) (optional)
