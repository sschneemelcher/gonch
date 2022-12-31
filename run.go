package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func runCommand(words []string) (string, string) {
	var cmd *exec.Cmd

	// Depending on the amount of words, we might need to use args
	if len(words) == 1 {
		cmd = exec.Command(words[0])
	} else if len(words) > 1 {
		cmd = exec.Command(words[0], words[1:]...)
	} else {
		// If the string was empty, do nothing
		return "", ""
	}

	// Check for builtins
	if words[0] == "exit" {
		os.Exit(0)
	}

	// Setup pipes so we can get the output of the command
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		os.Exit(1)
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		os.Exit(1)
	}

	err = cmd.Start()
	if err != nil {
		// If we get an error from exec, we want to have this in stderr too
		return "", err.Error()
	}

	// Read the output of the command
	stdout, err := ioutil.ReadAll(stdoutPipe)
	if err != nil {
		os.Exit(1)
		// fmt.Println(err)
	}
	stderr, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		// This prints `Exit status 1` etc..
		// fmt.Println(err)
	}

	return string(stdout), string(stderr)
}
