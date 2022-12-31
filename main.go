package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func prompt(r *bufio.Reader) []string {
	fmt.Print("$ ")
	line, err := r.ReadString('\n')
	if err != nil {
		fmt.Println("Cannot read from stdin")
		os.Exit(1)
	}
	line = strings.TrimSpace(line)
	words := strings.Fields(line)

	return words
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for true {

		// Print prompt, read line from user input and split it into words
		words := prompt(reader)

		var cmd *exec.Cmd

		// Depending on the amount of words, we might need to use args
		if len(words) == 1 {
			cmd = exec.Command(words[0])
		} else if len(words) > 1 {
			cmd = exec.Command(words[0], words[1:]...)
		} else {
			continue
		}

		// Check for builtins
		if words[0] == "exit" {
			os.Exit(0)
		}

		// Setup pipes so we can get the output of the command
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err)
		}

		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println(err)
		}

		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
		}

		// Read the output of the command
		stdout, err := ioutil.ReadAll(stdoutPipe)
		if err != nil {
			fmt.Println(err)
		}
		stderr, err := ioutil.ReadAll(stderrPipe)
		if err != nil {
			fmt.Println(err)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
		}

		// Print command output
		if len(stdout) > 0 {
			fmt.Print(string(stdout))
		}
		if len(stderr) > 0 {
			fmt.Print(string(stderr))
		}
	}
}
