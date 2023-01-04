package main

import (
	"bufio"
	"fmt"
	"os"
)

func printPrompt() {
	prompt := os.Getenv("PS1")
	fmt.Print(prompt)
}

func readInput(r *bufio.Reader) string {

	line, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Cannot read from stdin")
		os.Exit(1)
	}

	return line
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for true {

		// Print prompt, read line from user input and split it into words
		printPrompt()
		input := readInput(reader)
		words := parseInput(input)

		stdout, stderr := runCommand(words)

		// Print command output
		if len(stdout) > 0 {
			fmt.Print(stdout)
		}
		if len(stderr) > 0 {
			fmt.Println(stderr)
		}

	}
}
