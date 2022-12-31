package main

import (
	"bufio"
	"fmt"
	"os"
)

func prompt(r *bufio.Reader) []string {
	fmt.Print("$ ")
	line, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Cannot read from stdin")
		os.Exit(1)
	}

	return parseInput(line)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {

		// Print prompt, read line from user input and split it into words
		words := prompt(reader)

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
