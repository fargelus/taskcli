package main

import (
	"fmt"
	"os"
	"regexp"
)

var commands = [...]string{"add", "list", "delete", "done"}

type Task struct {
	id   int
	name string
}

func checkCommandWithArg(cmd string, arg string) {
	if !isValidCommand(cmd) {
		fmt.Fprintf(os.Stderr, "Error: Unknown command\n")
		os.Exit(1)
	}

	if !isCommandArgExpected(cmd, arg) {
		fmt.Fprintf(os.Stderr, "Error: Unknown arg for command <%s>\n", cmd)
		os.Exit(1)
	}
}

func isValidCommand(cmd string) bool {
	valid := false
	for _, item := range commands {
		if item == cmd {
			valid = true
			break
		}
	}

	return valid
}

func isCommandArgExpected(cmd string, arg string) bool {
	switch cmd {
	case "add":
		phraseRe := regexp.MustCompile(`\A\w+\z`)
		return phraseRe.MatchString(arg)
	case "done", "delete":
		numRe := regexp.MustCompile(`\A\d+\z`)
		return numRe.MatchString(arg)
	case "list":
		return arg == ""
	default:
		return false
	}
}

// go run main.go add "Купить хлеб"
// go run main.go list
// go run main.go done 1
// go run main.go delete 1

func main() {
	cmd := os.Args[1]
	arg := ""
	if len(os.Args) > 2 {
		arg = os.Args[2]
	}
	checkCommandWithArg(cmd, arg)

	if cmd == "add" {
		task := Task{id: 1, name: arg}
		fmt.Println("Task is", task)
	}
}
