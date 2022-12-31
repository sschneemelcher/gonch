package main

import "testing"

func TestSimpleCommand(t *testing.T) {
	// Set up the test by running the command and capturing the output
	input := "echo hello world"
	words := parseInput(input)

	stdout, _ := runCommand(words)

	// Check the output to make sure it matches the expected output
	expectedOutput := "hello world\n"
	if stdout != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, stdout)
	}
}