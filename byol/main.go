package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

// PROMPT is the message in the repl
const PROMPT = "lispy> "

func start() error {
	rl, err := readline.New(PROMPT)
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		fmt.Println(line)
	}
	return nil
}

func main() {
	fmt.Println("Lispy Version 0.0.1")
	fmt.Println("Press Ctrl+c to Exit")

	start()
}
