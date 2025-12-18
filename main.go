package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: teste-cli-go command [arguments]")
		fmt.Println("Available commands: list, read, create, update, delete")
		return
	}
	switch os.Args[1] {
	case "list":
		ListIssues(os.Args[2:])
	default:
		fmt.Println("unknown command")
	}
}
