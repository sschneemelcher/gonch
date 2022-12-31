package main

import "testing"

func TestEmptyCommand(t *testing.T) {
	// Set up the test by running the command and capturing the output
	input := ""
	words := parseInput(input)

	stdout, stderr := runCommand(words)

	// Check the output to make sure it matches the expected output
	expectedOutput := ""
	if stdout != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, stdout)
	}

	// Check the output to make sure it matches the expected output
	expectedOutput = ""
	if stderr != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, stderr)
	}
}

func TestSimpleCommandStderr(t *testing.T) {
	// Set up the test by running the command and capturing the output
	input := "man"
	words := parseInput(input)

	_, stderr := runCommand(words)

	// Check the output to make sure it matches the expected output
	expectedOutput := "What manual page do you want?\n"
	if stderr != expectedOutput {
		t.Errorf("Unexpected output. Expected: %q, Got: %q", expectedOutput, stderr)
	}
}

func TestSimpleCommandWithArgs(t *testing.T) {
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
